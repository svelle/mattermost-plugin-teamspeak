import PropTypes from 'prop-types';

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
