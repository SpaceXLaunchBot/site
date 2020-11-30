import React, { useState, useEffect } from 'react';
import './App.sass';
import Launch from './Launch';
import GetNextLaunch from './SpaceX';

export default function App() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [launchInfo, setLaunchInfo] = useState({});

    useEffect(() => {
        GetNextLaunch()
            .then((json) => {
                if (json === {}) {
                    setError({ message: 'SpaceX API request failed' });
                } else {
                    setLaunchInfo(json);
                }
                setIsLoaded(true);
            });
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

    return [
        <Launch launchInfo={launchInfo} />,
    ];
}
