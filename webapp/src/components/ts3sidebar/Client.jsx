import PropTypes from 'prop-types';

const Client = ({info}) => {
    return (
        <div
            className='ts3app-client'
            data-deafened={
                info.client_output_muted || !info.client_output_hardware
            }
            data-muted={
                info.client_input_muted ||
                !info.client_input_hardware ||
                info.client_away
            }
        >
            <span>{info.client_nickname}</span>
            <i
                className={`icon fa fa-${(
                    info.client_platform || ''
                ).toLowerCase()}`}
            />
        </div>
    );
};

Client.propTypes = {
    info: PropTypes.object.isRequired,
};

export default Client;
