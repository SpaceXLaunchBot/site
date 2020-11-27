import React, { useState, useEffect } from 'react';
import Countdown from './Countdown';

export default function Launch() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [nextLaunch, setNextLaunch] = useState({});

    useEffect(() => {
        fetch('https://api.spacexdata.com/v4/launches/next')
            .then((res) => res.json())
            .then(
                (json) => {
                    setNextLaunch(json);
                    setIsLoaded(true);
                },
                (e) => {
                    setIsLoaded(true);
                    setError(e);
                },
            );
    }, []);

    if (error) {
        return (
            <div>
                {`Error: ${error.message}`}
            </div>
        );
    } if (!isLoaded) {
        return <div>Loading...</div>;
    }

    let img;
    if (nextLaunch.links.patch.small !== null) {
        img = <img src={nextLaunch.links.patch.small} alt={`${nextLaunch.name} mission patch`} />;
    }

    return (
        <div>
            <h1>{nextLaunch.name}</h1>
            {img}
            <div className="inline">
                <p>Launching in </p>
                <Countdown futureDate={Date.parse(nextLaunch.date_utc)} />
            </div>
            <p>{nextLaunch.details}</p>
        </div>
    );
}
