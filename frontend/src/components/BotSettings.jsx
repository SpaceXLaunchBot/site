import React, { useEffect, useState } from 'react';
import { Box } from '@material-ui/core';
import Login from './Login';
import Channel from './Channel';
import Guild from './Guild';
import getSubscribed from '../internalapi/subscribed';

function monthDiff(dateFrom, dateTo) {
  // https://stackoverflow.com/a/4312956/6396652
  return dateTo.getMonth() - dateFrom.getMonth()
    + (12 * (dateTo.getFullYear() - dateFrom.getFullYear()));
}

export default function BotSettings() {
  const [discordOAuthToken, setDiscordOAuthToken] = useState('');
  const [loggedIn, setLoggedIn] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [subscribedInfo, setSubscribedInfo] = useState({});

  useEffect(async () => {
    let storedToken = localStorage.getItem('discord-oauth-token');
    const storedLoginTime = localStorage.getItem('discord-login-time');

    if (storedLoginTime !== null) {
      const now = new Date();
      const before = new Date(parseInt(storedLoginTime, 10));
      if (monthDiff(now, before) >= 1) {
        localStorage.removeItem('discord-oauth-token');
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
        setDiscordOAuthToken(storedToken);
      }
    }

    // We either found a token in localstorage or the url.
    if (storedToken !== null) {
      setLoggedIn(true);
      // Effects run asynchronously away from the actual render, so this will re-render
      // when setLoggedIn gets called above and the below web request will still be
      // happening in the background.
      const json = await getSubscribed(storedToken);
      setSubscribedInfo(json);
      setLoaded(true);
    }
  }, []);

  if (!loggedIn) {
    return <Login />;
  }
  if (!loaded) {
    return <p>Loading subscribed channel info...</p>;
  }
  if (subscribedInfo.success === false) {
    return <p>{`Error: ${subscribedInfo.error}`}</p>;
  }

  // The IDs are used as keys, just because they are there and are unique.
  const guilds = [];

  // TODO: fix
  // eslint-disable-next-line guard-for-in
  for (const guildId in subscribedInfo) {
    const subbedChannelsElems = [];
    for (const channel of subscribedInfo[guildId].subscribed_channels) {
      subbedChannelsElems.push(
        <Channel
          key={channel.id}
          info={channel}
          guildId={guildId}
          discordOAuthToken={discordOAuthToken}
        />,
      );
    }

    const guildInfo = subscribedInfo[guildId];
    guilds.push(
      <Guild key={guildId} name={guildInfo.name} icon={guildInfo.icon}>
        {subbedChannelsElems}
      </Guild>,
    );
  }

  return (
    <Box>
      {guilds}
    </Box>
  );
}
