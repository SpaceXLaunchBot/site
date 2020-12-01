import React, { useEffect, useState } from 'react';
import GetGuildsWithSubscribed from './Api';

const oauthUrl = 'https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=http%3A%2F%2Flocalhost%3A8080&response_type=token&scope=identify%20guilds';

export default function Login() {
    const [loading, setLoading] = useState(false);
    const [jsonData, setJsonData] = useState(undefined);

    useEffect(() => {
        const fragment = new URLSearchParams(window.location.hash.slice(1));
        if (fragment.has('access_token')) {
            setLoading(true);
            window.history.pushState(null, document.title, '/');
            const token = fragment.get('access_token');
            GetGuildsWithSubscribed(token)
                .then((json) => setJsonData(json))
                .then(() => { setLoading(false); });
        }
    }, []);

    if (loading === true) {
        return (
            <p>Loading...</p>
        );
    }
    if (jsonData === undefined) {
        return (
            <a href={oauthUrl}>Login with Discord</a>
        );
    }

    const info = [];
    jsonData.forEach((guild, i) => {
        info.push(<p key={i}>{guild.name}</p>);
    });

    return (
        <div>{info}</div>
    );
}
