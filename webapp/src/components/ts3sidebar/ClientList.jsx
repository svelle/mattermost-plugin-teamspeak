// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import PropTypes from 'prop-types';
import React from 'react';

import Client from './Client';

const ClientList = ({clients}) => {
    if (!Array.isArray(clients)) {
        return null;
    }
    return (
        <div className='ts3app-clientlist'>
            {clients.map((client) => (
                <Client
                    key={client.clid}
                    info={client}
                />
            ))}
        </div>
    );
};

ClientList.propTypes = {
    clients: PropTypes.array.isRequired,
};

export default ClientList;
