// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import manifest from '../../manifest';

const ChannelHeaderButton = () => (
    <img
        src={`/plugins/${manifest.id}/public/teamspeak5.svg`}
        className='ts3app-icon'
    />
);

export default ChannelHeaderButton;

