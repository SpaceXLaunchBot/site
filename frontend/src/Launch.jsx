import React, { useState, useEffect } from 'react';

export default function Launch() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [nextLaunch, setNextLaunch] = useState({});
    // const [timeToLaunch, setTimeToLaunch] = useState(0);

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
                Error:
                {error.message}
            </div>
        );
    } if (!isLoaded) {
        return <div>Loading...</div>;
    }
    return (
        <div>
            <h1>{nextLaunch.name}</h1>
            <img src={nextLaunch.links.patch.small} alt={`${nextLaunch.name} mission patch`} />
            <h2>{nextLaunch.details}</h2>
        </div>
    );
}
