import React, { useEffect, useState } from 'react';
import { Box } from '@material-ui/core';
import Login from './login';
import Channel from './channel';
import Guild from './guild';
import GetGuildsWithSubscribed from '../internalapi/guildswithsubscribed';

export default function BotSettings() {
    const [loggedIn, setLoggedIn] = useState(false);
    const [loaded, setLoaded] = useState(false);
    const [subscribedInfo, setSubscribedInfo] = useState({});

    useEffect(() => {
        const fragment = new URLSearchParams(window.location.hash.slice(1));
        if (fragment.has('access_token')) {
            window.history.pushState(null, document.title, '/');
            const accessToken = fragment.get('access_token');
            setLoggedIn(true);
            // Effects run asynchronously away from the actual render, so this will re-render when
            // setLoggedIn gets called above and the below web request will still be happening in
            // the background.
            GetGuildsWithSubscribed(accessToken)
                .then((json) => {
                    setSubscribedInfo(json);
                    setLoaded(true);
                });
        }
    }, []);

    if (!loggedIn) {
        return <Login />;
    }
    if (!loaded) {
        return <p>Loading!</p>;
    }

    // The IDs are used as keys, just because they are there and are unique.
    // TODO: What happens if API returns error or user has no servers?
    const guilds = [];
    for (const guildId in subscribedInfo) {
        const subbedChannelsElems = [];
        for (const channel of subscribedInfo[guildId].subscribed_channels) {
            subbedChannelsElems.push(<Channel key={channel.id} info={channel} />);
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
