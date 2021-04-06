import React from 'react';

type Props = {
    shouldHighlight: boolean,
};

export default function ChannelHeaderButton(props: Props) {
    let btnClass = 'icon fa fa-teamspeak';
    if (props.shouldHighlight) {
        btnClass += 'ts3-plugin-icon--activate';
    }

    return (
        <i className={btnClass}/>
    );
}

