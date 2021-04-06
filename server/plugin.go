package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/mattermost/mattermost-server/v5/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration  *configuration
	channels       []ts3ChannelInfo
	clients        map[int][]ts3ClientInfo
	cancelBgWorker context.CancelFunc
}

type ts3ClientListOptions struct {
	UID      bool `json:"-uid,omitempty,string"`
	Away     bool `json:"-away,omitempty,string"`
	Voice    bool `json:"-voice,omitempty,string"`
	Times    bool `json:"-times,omitempty,string"`
	Groups   bool `json:"-groups,omitempty,string"`
	Info     bool `json:"-info,omitempty,string"`
	Country  bool `json:"-country,omitempty,string"`
	Ip       bool `json:"-ip,omitempty,string"`
	Icon     bool `json:"-icon,omitempty,string"`
	Badges   bool `json:"-badges,omitempty,string"`
	Location bool `json:"-location,omitempty,string"`
}

type ts3ChannelListOptions struct {
	Topic        bool `json:"-topic,omitempty,string"`
	Flags        bool `json:"-flags,omitempty,string"`
	Voice        bool `json:"-voice,omitempty,string"`
	Limits       bool `json:"-limits,omitempty,string"`
	SecondsEmpty bool `json:"-secondsempty,omitempty,string"`
	Banners      bool `json:"-banners,omitempty,string"`
	Icon         bool `json:"-icon,omitempty,string"`
}

type ts3StatusResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ts3Boolean bool

func (t *ts3Boolean) UnmarshalJSON(b []byte) (err error) {
	var text string
	if err := json.Unmarshal(b, &text); err != nil {
		return err
	}
	if v, err := strconv.ParseBool(text); err != nil {
		return err
	} else {
		*t = ts3Boolean(v)
	}
	return nil
}

type ts3ClientInfo struct {
	CID                                  int        `json:"cid,omitempty,string"`
	CLID                                 int        `json:"clid,omitempty,string"`
	ClientAway                           ts3Boolean `json:"client_away,omitempty"`
	ClientAwayMessage                    string     `json:"client_away_message,omitempty"`
	ClientBadges                         string     `json:"client_badges,omitempty"`
	ClientChannelGroupId                 int        `json:"client_channel_group_id,omitempty,string"`
	ClientChannelGroupInheritedChannelId int        `json:"client_channel_group_inherited_channel_id,omitempty,string"`
	ClientCountry                        string     `json:"client_country,omitempty"`
	ClientCreated                        int        `json:"client_created,omitempty,string"`
	ClientDatabase_id                    int        `json:"client_database_id,omitempty,string"`
	ClientEstimatedLocation              string     `json:"client_estimated_location,omitempty"`
	ClientFlagTalking                    int        `json:"client_flag_talking,omitempty,string"`
	ClientIconId                         int        `json:"client_icon_id,omitempty,string"`
	ClientIdleTime                       int        `json:"client_idle_time,omitempty,string"`
	ClientInputHardware                  ts3Boolean `json:"client_input_hardware,omitempty"`
	ClientInputMuted                     ts3Boolean `json:"client_input_muted,omitempty"`
	ClientIsChannelCommander             ts3Boolean `json:"client_is_channel_commander,omitempty"`
	ClientIsPrioritySpeaker              ts3Boolean `json:"client_is_priority_speaker,omitempty"`
	ClientIsRecording                    ts3Boolean `json:"client_is_recording,omitempty"`
	ClientIsTalker                       ts3Boolean `json:"client_is_talker,omitempty"`
	ClientLastConnected                  int        `json:"client_lastconnected,omitempty,string"`
	ClientNickname                       string     `json:"client_nickname,omitempty"`
	ClientOutputHardware                 ts3Boolean `json:"client_output_hardware,omitempty"`
	ClientOutputMuted                    ts3Boolean `json:"client_output_muted,omitempty"`
	ClientPlatform                       string     `json:"client_platform,omitempty"`
	ClientServerGroups                   string     `json:"client_servergroups,omitempty"`
	ClientTalkPower                      int        `json:"client_talk_power,omitempty,string"`
	ClientType                           int        `json:"client_type,omitempty,string"`
	ClientUniqueIdentifier               string     `json:"client_unique_identifier,omitempty"`
	ClientVersion                        string     `json:"client_version,omitempty"`
	ConnectionClientIp                   string     `json:"connection_client_ip,omitempty"`
}

type ts3ChannelInfo struct {
	CID                         int        `json:"cid,omitempty,string"`
	ChannelBannerMode           int        `json:"channel_banner_mode,omitempty,string"`
	ChannelBanner_Graphic_URL   string     `json:"channel_banner_gfx_url,omitempty"`
	ChannelCodec                int        `json:"channel_codec,omitempty,string"`
	ChannelCodecQuality         int        `json:"channel_codec_quality,omitempty,string"`
	ChannelFlagDefault          ts3Boolean `json:"channel_flag_default,omitempty"`
	ChannelFlagPassword         ts3Boolean `json:"channel_flag_password,omitempty"`
	ChannelFlagPermanent        ts3Boolean `json:"channel_flag_permanent,omitempty"`
	ChannelFlagSemiPermanent    ts3Boolean `json:"channel_flag_semi_permanent,omitempty"`
	ChannelIconId               int        `json:"channel_icon_id,omitempty,string"`
	ChannelMaxClients           int        `json:"channel_maxclients,omitempty,string"`
	ChannelMaxFamilyClients     int        `json:"channel_maxfamilyclients,omitempty,string"`
	ChannelName                 string     `json:"channel_name,omitempty"`
	ChannelNeededSubscribePower int        `json:"channel_needed_subscribe_power,omitempty,string"`
	ChannelNeededTalkPower      int        `json:"channel_needed_talk_power,omitempty,string"`
	ChannelOrder                int        `json:"channel_order,omitempty,string"`
	ChannelTopic                string     `json:"channel_topic,omitempty"`
	PID                         int        `json:"pid,omitempty,string"`
	SecondsEmpty                int        `json:"seconds_empty,omitempty,string"`
	TotalClients                int        `json:"total_clients,omitempty,string"`
	TotalClientsFamily          int        `json:"total_clients_family,omitempty,string"`
}

type ts3ChannelListResponse struct {
	Status ts3StatusResponse `json:"status"`
	Body   []ts3ChannelInfo  `json:"body"`
}

type ts3ClientListResponse struct {
	Status ts3StatusResponse `json:"status"`
	Body   []ts3ClientInfo   `json:"body"`
}

func (p *Plugin) updateData(ctx context.Context, time <-chan time.Time) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time:
			if err := p.UpdateChannelList(); err != nil {
				p.API.LogDebug(err.Error())
				continue
			}
			if err := p.UpdateClientList(); err != nil {
				p.API.LogDebug(err.Error())
				continue
			}
		}
	}
}

func (p *Plugin) OnActivate() (err error) {
	ticker := time.NewTicker(30 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	p.cancelBgWorker = cancel
	go p.updateData(ctx, ticker.C)
	return nil
}

func (p *Plugin) OnDeactivate() error {
	p.cancelBgWorker()
	return nil
}

func (p *Plugin) UpdateChannelList() (err error) {
	config := p.getConfiguration()
	channelListOptions := ts3ChannelListOptions{true, true, true, true, true, true, true}
	channelListURL := fmt.Sprintf("%s/%d/channellist", config.WebQueryURL, config.ServerId)
	client := http.Client{}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	if err := encoder.Encode(channelListOptions); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", channelListURL, &buffer)
	if err != nil {
		return err
	}
	req.Header.Add("x-api-key", config.APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code %d returned: %v", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var channelList ts3ChannelListResponse
	if err := decoder.Decode(&channelList); err != nil {
		return err
	}
	if channelList.Status.Code == 0 {

		p.channels = channelList.Body
		sort.Sort(ts3ChannelByOrderId(p.channels))
		return nil
	}
	return fmt.Errorf(channelList.Status.Message)
}

func (p *Plugin) UpdateClientList() (err error) {
	config := p.getConfiguration()
	clientListOptions := ts3ClientListOptions{true, true, true, true, true, true, true, true, true, true, true}
	clientListURL := fmt.Sprintf("%s/%d/clientlist", config.WebQueryURL, config.ServerId)
	client := http.Client{}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	if err := encoder.Encode(clientListOptions); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", clientListURL, &buffer)
	if err != nil {
		return err
	}
	req.Header.Add("x-api-key", config.APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("invalid status code %d returned: %v", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var clientList ts3ClientListResponse
	if err := decoder.Decode(&clientList); err != nil {
		return err
	}
	if clientList.Status.Code == 0 {
		p.clients = make(map[int][]ts3ClientInfo)
		for _, client := range clientList.Body {
			p.clients[client.CID] = append(p.clients[client.CID], client)
		}
		return nil
	}
	return fmt.Errorf(clientList.Status.Message)
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("Mattermost-User-Id")
	if userId == "" {
		w.WriteHeader(403)
		return
	}

	encoder := json.NewEncoder(w)

	url := fmt.Sprintf("/plugins/%s/info", manifest.Id)
	if r.RequestURI != url {
		w.WriteHeader(404)
		return
	}
	w.WriteHeader(200)
	encoder.Encode(struct {
		Channels []ts3ChannelInfo
		Clients  map[int][]ts3ClientInfo
	}{
		Channels: p.channels,
		Clients:  p.clients,
	})
}

type ts3ChannelByOrderId []ts3ChannelInfo

// Len is the number of elements in the collection.
func (t ts3ChannelByOrderId) Len() int {
	return len(t)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (t ts3ChannelByOrderId) Less(i int, j int) bool {
	return t[i].ChannelOrder < t[j].ChannelOrder
}

// Swap swaps the elements with indexes i and j.
func (t ts3ChannelByOrderId) Swap(i int, j int) {
	t[i], t[j] = t[j], t[i]
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
