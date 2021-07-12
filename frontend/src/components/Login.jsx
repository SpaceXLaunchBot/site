import React from 'react';
import { Button } from '@material-ui/core';

const loc = encodeURI(`${window.location.protocol}//${window.location.host}`);
const oauthUrl = `https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=${loc}&response_type=token&scope=identify%20guilds`;

export default function Login(props) {
  const { loggedIn, logOut } = props;

  if (!loggedIn) {
    return (
      <Button
        variant="contained"
        color="secondary"
        // eslint-disable-next-line no-return-assign
        onClick={() => window.location.href = oauthUrl}
      >
        Login
      </Button>
    );
  }

  return (
    <Button
      variant="contained"
      color="secondary"
      className="logOutBtn"
      onClick={logOut}
    >
      Log Out
    </Button>
  );
}
