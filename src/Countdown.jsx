import React, { useState, useEffect } from 'react';
import moment from 'moment';

export default function Countdown(props) {
    const [diffDuration, setDiffDuration] = useState(undefined);
    const { futureMoment } = props;

    useEffect(() => {
        setDiffDuration(moment.duration(futureMoment.diff(moment())));

        const loop = setInterval(() => {
            setDiffDuration(moment.duration(futureMoment.diff(moment())));
        }, 1000);

        return function cleanup() {
            clearInterval(loop);
        };
    }, [futureMoment]);

    if (diffDuration === undefined) {
        return <p />;
    }
    return (
        <p>
            {`${diffDuration.days()} days, ${diffDuration.hours()} hours, ${diffDuration.minutes()} minutes, ${diffDuration.seconds()} seconds`}
        </p>
    );
}
