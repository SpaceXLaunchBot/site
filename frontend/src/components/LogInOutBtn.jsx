import React from 'react';
import { Button, Icon } from '@material-ui/core';
import '../css/Login.scss';

const loginRedirect = encodeURI(`${window.location.protocol}//${window.location.host}/login`);
const oauthUrl = `https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=${loginRedirect}&response_type=code&scope=identify%20guilds`;

export default function LogInOutBtn(props) {
  const { loggedIn, setLoggedIn } = props;

  const login = () => {
    window.location.href = oauthUrl;
  };

  const logout = async () => {
    const res = await fetch('/api/auth/logout');
    const json = await res.json();
    if (json.success === true) {
      localStorage.removeItem('user-data');
      setLoggedIn(false);
    } else {
      // TODO: Give a toast popup or something.
    }
  };

  if (!loggedIn) {
    return <Button variant="contained" onClick={login}>Login</Button>;
  }

  return (
    <div className="logOut" onClick={logout} onKeyDown={logout} role="button" tabIndex={0}>
      <Icon className="logOutIcon" alt="Log Out" title="Log Out">exit_to_app</Icon>
    </div>
  );
}
