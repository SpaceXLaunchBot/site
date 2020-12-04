import React, { useEffect, useState } from 'react';
import './app.sass';
import { makeStyles, ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import { Grid } from '@material-ui/core';
import { ToastProvider } from 'react-toast-notifications';
import Launch from './components/launch';
import GetNextLaunch from './spacexapi/nextlaunch';
import theme from './theme';
import BotSettings from './components/botsettings';

const useStyles = makeStyles(() => ({
    root: {
        textAlign: 'center',
    },
}));

export default function App() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [launchInfo, setLaunchInfo] = useState({});

    const classes = useStyles();

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
        // TODO: This error and loading should just be for launch, not everything.
        //  To do this, probably also a good idea to move the loading effect into launch.
        return (
            <div>
                {`Error: ${error.message}`}
            </div>
        );
    } if (!isLoaded) {
        return <div>Loading...</div>;
    }

    return (
        <ToastProvider autoDismiss>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                <Grid
                    container
                    direction="row"
                    justify="space-evenly"
                    alignItems="stretch"
                    classes={classes}
                >
                    {/* Both take up half a width on larger screens or full width on smaller */}
                    <Grid item xs={12} sm={6}>
                        <Launch launchInfo={launchInfo} />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <BotSettings />
                    </Grid>
                </Grid>
            </ThemeProvider>
        </ToastProvider>
    );
}
