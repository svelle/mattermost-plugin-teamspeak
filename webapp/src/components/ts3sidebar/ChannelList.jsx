import {useEffect, useState} from 'react';

import {id as pluginId} from '../../manifest';

import Channel from './Channel';

function sort_channels(channels, clients) {
    const channel_map = new Map(
        channels.map((channel) => [channel.channel_order || 0, channel]),
    );
    const sorted_channels = [];
    for (
        let channel = channel_map.get(0);
        channel != null;
        channel = channel_map.get(channel.cid)
    ) {
        channel.clients = (clients[channel.cid] || []).filter(
            (client) => client.client_type !== '1',
        );
        if (channel.clients.length > 0) {
            sorted_channels.push(channel);
        }
    }
    return sorted_channels;
}

const ChannelList = () => {
    const [channels, setChannels] = useState(null);
    useEffect(() => {
        (async () => {
            const resp = await fetch(`/plugins/${pluginId}/info`);
            if (!resp.ok) {
                return;
            }
            const data = await resp.json();
            if (data.Channels == null) {
                return;
            }
            setChannels(sort_channels(data.Channels, data.Clients));
        })();
    }, [setChannels]);
    if (channels == null) {
        return null;
    }
    return (
        <div className='ts3app'>
            <div className='ts3app-channellist'>
                {channels.map((channel) => (
                    <Channel
                        key={channel.cid}
                        info={channel}
                    />
                ))}
            </div>
        </div>
    );
};

export default ChannelList;
