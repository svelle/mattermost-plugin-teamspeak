// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useEffect, useState} from 'react';

import Channel from './Channel';

import manifest from '../../manifest';

function sortChannels(channels, clients) {
    const channelMap = new Map(
        channels.map((channel) => [channel.channel_order || 0, channel]),
    );
    const sortedChannels = [];
    for (
        let channel = channelMap.get(0);
        channel != null;
        channel = channelMap.get(channel.cid)
    ) {
        channel.clients = (clients[channel.cid] || []).filter(
            (client) => client.client_type !== '1',
        );
        if (channel.clients.length > 0) {
            sortedChannels.push(channel);
        }
    }
    return sortedChannels;
}

const ChannelList = () => {
    const [channels, setChannels] = useState(null);
    useEffect(() => {
        (async () => {
            const resp = await fetch(`/plugins/${manifest.id}/info`);
            if (!resp.ok) {
                return;
            }
            const data = await resp.json();
            if (data.Channels == null) {
                return;
            }
            setChannels(sortChannels(data.Channels, data.Clients));
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
