import React, { useEffect, useState } from 'react';
import Loader from 'react-loader-spinner';
import Channel from '../components/Channel';
import Guild from '../components/Guild';
import getSubscribed from '../internalapi/subscribed';
import '../css/Settings.scss';

export default function Settings(props) {
  const { discordOAuthToken, loggedIn } = props;
  const [loaded, setLoaded] = useState(false);
  const [error, setError] = useState('');
  const [subscribedInfo, setSubscribedInfo] = useState({});

  useEffect(async () => {
    if (loggedIn) {
      try {
        const json = await getSubscribed(discordOAuthToken);
        setSubscribedInfo(json);
      } catch (e) {
        // https://davidwalsh.name/detect-error-type-javascript
        if (e.constructor === SyntaxError) {
          setError('Server returned invalid JSON');
        } else {
          setError(e.toString());
        }
      }
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

  if (error !== '' || subscribedInfo.success === false) {
    return (
      <div>
        <h2>Failed to get data</h2>
        <p>{error !== '' ? error : subscribedInfo.error}</p>
      </div>
    );
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
    <div>
      {guilds}
    </div>
  );
}
