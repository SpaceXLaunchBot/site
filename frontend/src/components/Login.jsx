import React from 'react';
import { Button } from '@material-ui/core';
import '../css/Login.scss';

const loc = encodeURI(`${window.location.protocol}//${window.location.host}`);
const oauthUrl = `https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=${loc}&response_type=token&scope=identify%20guilds`;

export default function Login(props) {
  const { loggedIn, logOut } = props;

  const setRed = (e) => {
    e.target.setAttribute('src', '/logout-red.svg');
  };

  const setPink = (e) => {
    e.target.setAttribute('src', '/logout-pink.svg');
  };

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
    <div className="logOut" onClick={logOut} onKeyDown={logOut} role="button" tabIndex={0}>
      <img
        className="logOutImg"
        src="/logout-pink.svg"
        onMouseOver={setRed}
        onMouseOut={setPink}
        onFocus={setRed}
        onBlur={setPink}
        alt="Log Out"
        title="Log Out"
      />
    </div>
  );
}
