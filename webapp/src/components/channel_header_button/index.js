// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {connect} from 'react-redux';

import ChannelHeaderButton from './channel_header_button';

function mapStateToProps() {
    return {
        shouldHighlight: false,
    };
}

function mapDispatchToProp() {
    return {};
}

export default connect(mapStateToProps, mapDispatchToProp)(ChannelHeaderButton);
