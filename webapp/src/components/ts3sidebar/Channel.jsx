import PropTypes from 'prop-types';

import ClientList from './ClientList';

const Channel = ({info}) => {
    const spacer = info.channel_name.match(
        /^\[[lrc*]?spacer\d*\]([-_~.,]{3}|.*)$/,
    );
    return (
        <div className='ts3app-channel-container'>
            <div className={spacer ? 'ts3app-spacer' : 'ts3app-channel'}>
                <span className='ts3app-title'>
                    {spacer ? spacer[1] : info.channel_name}
                </span>
                {info.channel_topic && (
                    <span className='ts3app-tooltiptext'>
                        {info.channel_topic}
                    </span>
                )}
            </div>
            <ClientList clients={info.clients}/>
        </div>
    );
};

Channel.propTypes = {
    info: PropTypes.object.isRequired,
};

export default Channel;
