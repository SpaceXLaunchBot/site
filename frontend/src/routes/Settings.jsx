import React, { useEffect, useState } from 'react';
import Loader from '../components/Loader';
import ChannelSettings from '../components/ChannelSettings';
import getSubscribed from '../internalapi/subscribed';
import '../css/Settings.scss';

export default function Settings(props) {
  const { loggedIn } = { props };
  const [loaded, setLoaded] = useState(false);
  const [error, setError] = useState('');
  const [subscribedInfo, setSubscribedInfo] = useState({});

  useEffect(async () => {
    if (loggedIn) {
      try {
        const json = await getSubscribed();
        setSubscribedInfo(json);
      } catch (e) {
        setError(e.toString());
      }
      setLoaded(true);
    }
  }, [loggedIn]);

  if (!loggedIn) {
    return <h2>Login Required</h2>;
  }
  if (!loaded) {
    return <Loader />;
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
  const channels = [];

  // TODO: fix
  // eslint-disable-next-line guard-for-in
  for (const guildId in subscribedInfo.subscribed) {
    const guildInfo = subscribedInfo.subscribed[guildId];
    for (const channelInfo of guildInfo.subscribed_channels) {
      channels.push(
        <ChannelSettings
          key={channelInfo.id}
          guildId={guildId}
          guildName={guildInfo.name}
          guildIcon={guildInfo.icon}
          channelInfo={channelInfo}
        />,
      );
    }
  }

  return (
    <div className="settings">
      {channels}
    </div>
  );
}
