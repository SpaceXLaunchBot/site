import React, { useState, useEffect } from 'react';
import './app.sass';
import { ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Launch from './components/launch';
import GetNextLaunch from './spacexapi/nextlaunch';
import Login from './components/login';
import theme from './theme';
import BotSettings from './components/botsettings';

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

    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <Launch launchInfo={launchInfo} />
            <BotSettings />
        </ThemeProvider>
    );
}
