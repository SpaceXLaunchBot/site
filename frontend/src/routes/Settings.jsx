import React, { useEffect, useState } from 'react';
import { Box } from '@material-ui/core';
import Loader from 'react-loader-spinner';
import Channel from '../components/Channel';
import Guild from '../components/Guild';
import getSubscribed from '../internalapi/subscribed';
import '../css/Settings.scss';

export default function Settings(props) {
  const { discordOAuthToken, loggedIn } = props;
  const [loaded, setLoaded] = useState(false);
  const [subscribedInfo, setSubscribedInfo] = useState({});

  useEffect(async () => {
    if (loggedIn) {
      const json = await getSubscribed(discordOAuthToken);
      setSubscribedInfo(json);
      setLoaded(true);
    }
  }, [loggedIn]);

  if (!loggedIn) {
    return <h2>Login Required</h2>;
  }
  if (!loaded) {
    return (
      <Loader
        className="loader"
        type="Grid"
        color="#00BFFF"
        height={25}
        width={25}
      />
    );
  }
  if (subscribedInfo.success === false) {
    return <p>{`Error: ${subscribedInfo.error}`}</p>;
  }

  // The IDs are used as keys, just because they are there and are unique.
  const guilds = [];

  // TODO: fix
  // eslint-disable-next-line guard-for-in
  for (const guildId in subscribedInfo.subscribed) {
    const subbedChannelsElems = [];
    for (const channel of subscribedInfo.subscribed[guildId].subscribed_channels) {
      subbedChannelsElems.push(
        <Channel
          key={channel.id}
          info={channel}
          guildId={guildId}
          discordOAuthToken={discordOAuthToken}
        />,
      );
    }

    const guildInfo = subscribedInfo.subscribed[guildId];
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
