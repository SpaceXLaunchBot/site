import React from 'react';
import { Button, Icon } from '@material-ui/core';
import '../css/Login.scss';

const logoutPath = encodeURI(`${window.location.protocol}//${window.location.host}/api/logout`);
const loginRedirect = encodeURI(`${window.location.protocol}//${window.location.host}/api/login`);
const oauthUrl = `https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=${loginRedirect}&response_type=code&scope=identify%20guilds`;

export default function LogInOutBtn(props) {
  const { loggedIn } = props;

  const login = () => {
    window.location.href = oauthUrl;
  };

  const logout = () => {
    localStorage.removeItem('user-data');
    window.location.href = logoutPath;
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
