// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React, {useEffect, useState} from 'react';

import Channel from './Channel';

import manifest from '../../manifest';

function sortChannels(channels, clients) {
    const roots = {};
    const channelMap = new Map(channels.
        map((channel) => {
            channel.clients = (clients[channel.cid] || []).filter(
                (client) => client.client_type !== '1',
            );
            channel.children = {};
            channel.channel_order = Number(channel.channel_order) || 0;
            channel.cid = Number(channel.cid) || 0;
            channel.pid = Number(channel.pid) || 0;
            if (channel.pid === 0) {
                roots[channel.channel_order] = channel;
            }
            return [channel.cid, channel];
        }),
    );
    channelMap.values().forEach((channel) => {
        if (channel.pid !== 0) {
            channelMap.get(channel.pid).children[channel.channel_order] = channel;
        }
    });
    const channelList = [];
    const walker = (root) => {
        for (let channel = root[0]; channel != null; channel = root[channel.cid]) {
            channelList.push(channel);
            if (channel.children[0] != null) {
                walker(channel.children);
            }
        }
    };
    walker(roots);
    return channelList.filter((channel) => channel.clients.length > 0);
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
