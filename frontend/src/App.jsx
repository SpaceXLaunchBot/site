import React from 'react';
import './app.scss';
import CssBaseline from '@material-ui/core/CssBaseline';
import { Grid } from '@material-ui/core';
import { ToastProvider } from 'react-toast-notifications';
import Launch from './components/Launch';
import BotSettings from './components/BotSettings';

export default function App() {
  return (
    <ToastProvider autoDismiss>
      <CssBaseline />
      <Grid
        container
        direction="row"
        justify="space-evenly"
        alignItems="stretch"
      >
        <Grid item xs={12} sm={6}>
          <Launch />
        </Grid>
        <Grid item xs={12} sm={6}>
          <BotSettings />
        </Grid>
      </Grid>
    </ToastProvider>
  );
}
