import React, { useEffect, useState } from 'react';

const oauthUrl = 'https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=http%3A%2F%2Flocalhost%3A8080&response_type=token&scope=identify%20guilds';

export default function Login() {
    const [accessToken, setAccessToken] = useState('');

    useEffect(() => {
        const fragment = new URLSearchParams(window.location.hash.slice(1));
        if (fragment.has('access_token')) {
            setAccessToken(fragment.get('access_token'));
            window.history.pushState(null, document.title, '/');
        }
    }, []);

    if (accessToken === '') {
        return (
            <a href={oauthUrl}>Login with Discord</a>
        );
    }
    return (
        <p>{`You do be logged in doe: ${accessToken}`}</p>
    );
}
