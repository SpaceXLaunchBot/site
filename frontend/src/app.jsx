import React from 'react';
import './app.sass';
import { makeStyles, ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import { Grid } from '@material-ui/core';
import { ToastProvider } from 'react-toast-notifications';
import Launch from './components/launch';
import theme from './theme';
import BotSettings from './components/botsettings';

const useStyles = makeStyles(() => ({
    root: {
        textAlign: 'center',
    },
}));

export default function App() {
    const classes = useStyles();
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
                    <Grid item xs={12} sm={6}>
                        <Launch />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                        <BotSettings />
                    </Grid>
                </Grid>
            </ThemeProvider>
        </ToastProvider>
    );
}
