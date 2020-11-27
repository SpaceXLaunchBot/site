import React, { useState, useEffect } from 'react';

const dayMs = 1000 * 60 * 60 * 24;
const hourMs = 1000 * 60 * 60;
const minMs = 1000 * 60;
const secMs = 1000;

function howFarFromNow(futureDate) {
    const diffTime = futureDate - Date.now();
    return {
        days: Math.floor(diffTime / dayMs),
        hours: Math.floor(diffTime / hourMs) % 24,
        mins: Math.floor(diffTime / minMs) % 60,
        secs: Math.floor(diffTime / secMs) % 60,
    };
}

export default function Countdown(props) {
    const [time, setTime] = useState({
        days: 0, hours: 0, mins: 0, secs: 0,
    });

    useEffect(() => {
        // Do it immediately and also every second.
        setTime(howFarFromNow(props.futureDate));
        setInterval(() => {
            setTime(howFarFromNow(props.futureDate));
        }, 1000);
    }, []);

    return (
        <p>
            {`${time.days} days, ${time.hours} hours, ${time.mins} minutes, ${time.secs} seconds`}
        </p>
    );
}
