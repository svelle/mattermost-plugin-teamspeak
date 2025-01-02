// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import type {Action, Store} from 'redux';

import type {GlobalState} from 'mattermost-redux/types/store';

import ChannelHeaderButton from './components/channel_header_button/channel_header_button';
import ChannelList from './components/ts3sidebar/ChannelList';
import manifest from './manifest';
import './ts3Styles.css';
// eslint-disable-next-line import/no-unresolved
import type {PluginRegistry} from './types/mattermost-webapp';

export default class Plugin {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars, @typescript-eslint/no-empty-function
    public async initialize(registry: PluginRegistry, store: Store<GlobalState, Action<Record<string, unknown>>>) {
        // @see https://developers.mattermost.com/extend/plugins/webapp/reference/
        const {toggleRHSPlugin} = registry.registerRightHandSidebarComponent(ChannelList, 'TeamSpeak™ 3 Server');

        registry.registerChannelHeaderButtonAction(
            <ChannelHeaderButton/>,
            () => store.dispatch(toggleRHSPlugin),
            'TeamSpeak™ 3 Server',
            'Show people connected to our server',
        );
    }
}

declare global {
    interface Window {
        registerPlugin(id: string, plugin: Plugin): void;
    }
}

window.registerPlugin(manifest.id, new Plugin());
