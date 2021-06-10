import React, { useEffect, useState } from 'react';
import { Box } from '@material-ui/core';
import Login from './Login';
import Channel from './Channel';
import Guild from './Guild';
import getSubscribed from '../internalapi/subscribed';

export default function BotSettings() {
  const [discordOAuthToken, setDiscordOAuthToken] = useState(localStorage.getItem('discord-oauth-token') || '');
  const [loggedIn, setLoggedIn] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [subscribedInfo, setSubscribedInfo] = useState({});

  useEffect(async () => {
    // We use a local var here as the set state functions are async and don't change the state
    // variable straight away.
    let localToken = discordOAuthToken;

    if (localToken === '') {
      const fragment = new URLSearchParams(window.location.hash.slice(1));
      if (fragment.has('access_token')) {
        window.history.pushState(null, document.title, '/');
        localToken = fragment.get('access_token');
        localStorage.setItem('discord-oauth-token', localToken);
        setDiscordOAuthToken(localToken);
      }
    }

    // If we either already have or just discovered an access token.
    if (localToken !== '') {
      setLoggedIn(true);
      // Effects run asynchronously away from the actual render, so this will re-render
      // when setLoggedIn gets called above and the below web request will still be
      // happening in the background.
      const json = await getSubscribed(localToken);
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
