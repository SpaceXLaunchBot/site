import React, { useEffect } from 'react';
import { Button } from '@material-ui/core';

const oauthUrl = 'https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=http%3A%2F%2Flocalhost%3A8080&response_type=token&scope=identify%20guilds';

export default function Login(props) {
    const { setToken } = props;

    useEffect(() => {
        const fragment = new URLSearchParams(window.location.hash.slice(1));
        if (fragment.has('access_token')) {
            window.history.pushState(null, document.title, '/');
            const token = fragment.get('access_token');
            // TODO: Attempt to validate?
            setToken(token);
        }
    }, []);

    return (
        // eslint-disable-next-line no-return-assign
        <Button variant="contained" color="primary" onClick={() => window.location.href = oauthUrl}>Login with Discord</Button>
    );
}
