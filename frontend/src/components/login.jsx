import React from 'react';
import { Button } from '@material-ui/core';

let oauthUrl = 'https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=https%3A%2F%2Fspacexlaunchbot.simonj.tech%2F&response_type=token&scope=identify%20guilds';
if (window.location.hostname === 'localhost') {
    oauthUrl = 'https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=http%3A%2F%2Flocalhost%3A8080&response_type=token&scope=identify%20guilds';
}

export default function Login() {
    return (
        // eslint-disable-next-line no-return-assign
        <Button variant="contained" color="secondary" onClick={() => window.location.href = oauthUrl}>Login with Discord</Button>
    );
}
