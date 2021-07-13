import React, { useEffect, useState } from 'react';
import './css/App.scss';
import CssBaseline from '@material-ui/core/CssBaseline';
import { ToastProvider } from 'react-toast-notifications';
import { StylesProvider } from '@material-ui/core/styles';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import moment from 'moment';
import NavBar from './components/NavBar';
import Settings from './routes/Settings';
import Home from './routes/Home';
import Stats from './routes/Stats';
import Commands from './routes/Commands';

function isWithinAWeek(momentDate) {
  const aWeekAgo = moment().subtract(7, 'days').startOf('day');
  return momentDate.isAfter(aWeekAgo);
}

export default function App() {
  const [discordOAuthToken, setDiscordOAuthToken] = useState('');
  const [loggedIn, setLoggedIn] = useState(false);

  const logOut = () => {
    localStorage.removeItem('discord-oauth-token');
    localStorage.removeItem('discord-login-time');
    localStorage.removeItem('user-data');
    setDiscordOAuthToken('');
    setLoggedIn(false);
  };

  useEffect(async () => {
    let storedToken = localStorage.getItem('discord-oauth-token');
    const storedLoginTime = localStorage.getItem('discord-login-time');

    // Delete stored token if it's older than a week.
    if (storedLoginTime !== null) {
      if (!isWithinAWeek(moment(parseInt(storedLoginTime, 10)))) {
        logOut();
        storedToken = null;
      }
    }

    // Attempt to get token from url.
    if (storedToken === null) {
      const fragment = new URLSearchParams(window.location.hash.slice(1));
      if (fragment.has('access_token')) {
        window.history.pushState(null, document.title, '/');
        storedToken = fragment.get('access_token');
        localStorage.setItem('discord-oauth-token', storedToken);
        localStorage.setItem('discord-login-time', `${Date.now()}`);
      }
    }

    // We either found a token in localstorage or the url.
    if (storedToken !== null) {
      setDiscordOAuthToken(storedToken);
      setLoggedIn(true);
    }
  }, []);

  return (
    <ToastProvider autoDismiss>
      {/* See https://material-ui.com/guides/interoperability/#controlling-priority-2 */}
      <StylesProvider injectFirst>
        <CssBaseline />
        <BrowserRouter>
          <NavBar discordOAuthToken={discordOAuthToken} loggedIn={loggedIn} logOut={logOut} />
          <Switch>
            <Route path="/commands">
              <Commands />
            </Route>
            <Route path="/stats">
              <Stats />
            </Route>
            <Route path="/settings">
              <Settings discordOAuthToken={discordOAuthToken} loggedIn={loggedIn} />
            </Route>
            <Route path="/">
              <Home />
            </Route>
          </Switch>
        </BrowserRouter>
      </StylesProvider>
    </ToastProvider>
  );
}
