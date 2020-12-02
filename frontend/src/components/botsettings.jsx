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
        // TODO: Is localStorage insecure?
        // TODO: What happens if token is expired or incorrect?
        let accessToken = localStorage.getItem('discord-oauth-token') || '';
        if (accessToken === '') {
            const fragment = new URLSearchParams(window.location.hash.slice(1));
            if (fragment.has('access_token')) {
                window.history.pushState(null, document.title, '/');
                accessToken = fragment.get('access_token');
                localStorage.setItem('discord-oauth-token', accessToken);
            }
        }
        // looks like it should be else, but we want to this to trigger even if accessToken was ''.
        if (accessToken !== '') {
            setLoggedIn(true);
            // Effects run asynchronously away from the actual render, so this will re-render
            // when setLoggedIn gets called above and the below web request will still be
            // happening in the background.
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
        return <p>Loading subscribed channel info...</p>;
    }

    // The IDs are used as keys, just because they are there and are unique.
    // TODO: What happens if API returns error?
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
