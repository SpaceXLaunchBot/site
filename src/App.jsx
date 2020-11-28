import React, { useState, useEffect } from 'react';
import './App.sass';
import Launch from './Launch';
import NumberCarousel from './NumberCarousel';
import { GetLaunch, GetNextLaunch } from './SpaceX';

export default function App() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [launchInfo, setLaunchInfo] = useState({});
    const [launchNum, setLaunchNum] = useState(0);

    useEffect(() => {
        let dataPromise;
        if (launchNum > 0) {
            dataPromise = GetLaunch(launchNum);
        } else {
            dataPromise = GetNextLaunch();
        }
        dataPromise
            .then((json) => {
                if (json === {}) {
                    setError('API request failed');
                } else {
                    setLaunchInfo(json);
                }
                setIsLoaded(true);
            });
    }, [launchNum]);

    if (error) {
        return (
            <div>
                {`Error: ${error.message}`}
            </div>
        );
    } if (!isLoaded) {
        return <div>Loading...</div>;
    }

    return [
        <NumberCarousel number={launchInfo.flight_number} setNumber={setLaunchNum} />,
        <Launch launchInfo={launchInfo} />,
    ];
}
