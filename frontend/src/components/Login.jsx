import React from 'react';
import { Button } from '@material-ui/core';
import '../css/Login.scss';

const loc = encodeURI(`${window.location.protocol}//${window.location.host}`);
const oauthUrl = `https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=${loc}&response_type=token&scope=identify%20guilds`;

export default function Login(props) {
  const { loggedIn, logOut } = props;
  // TODO: Make smaller like avatar image?

  if (!loggedIn) {
    return (
      <Button
        variant="contained"
        // eslint-disable-next-line no-return-assign
        onClick={() => window.location.href = oauthUrl}
      >
        Login
      </Button>
    );
  }

  return (
    <div className="logOut" onClick={logOut} onKeyDown={logOut} role="button" tabIndex={0}>
      <img
        className="logOutImg"
        src="/logout.svg"
        alt="Log Out"
        title="Log Out"
      />
    </div>
  );
}
