import React from 'react';
import { Link } from 'react-router-dom';
import { Grid } from '@material-ui/core';
import Login from './Login';
import '../css/NavBar.scss';
import UserInfo from './UserInfo';

export default function NavBar(props) {
  const { discordOAuthToken, loggedIn, logOut } = props;
  return (
    <Grid
      container
      direction="row"
      justify="space-evenly"
      alignItems="stretch"
      className="navbar"
    >
      <Grid item xs={12} sm={6} className="navbarTitle">
        <img alt="SpaceXLaunchBot icon" src="/logo192.png" />
        <Link to="/"><h1>SpaceXLaunchBot</h1></Link>
      </Grid>
      <Grid item xs={12} sm={6} className="navbarLinks">
        <Link to="/settings">Bot Settings</Link>
        <Link to="/stats">Stats</Link>
        {/* No styling required, just having this adds enough space. */}
        <div className="spacer" />
        <UserInfo discordOAuthToken={discordOAuthToken} loggedIn={loggedIn} />
        <Login loggedIn={loggedIn} logOut={logOut} />
      </Grid>
    </Grid>
  );
}
