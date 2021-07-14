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
      <Grid item xs={12} sm={3} className="navbarTitle">
        <img className="circleImg" alt="SpaceXLaunchBot icon" src="/logo192.png" />
        <Link to="/"><h1>SpaceXLaunchBot</h1></Link>
      </Grid>
      <Grid item xs={12} sm={9} className="navbarLinks">
        <Link to="/commands">Commands</Link>
        <Link to="/settings">Server Settings</Link>
        <Link to="/stats">Stats</Link>
        <a href="https://top.gg/bot/411618411169447950/" rel="noreferrer" target="_blank">Top.gg</a>
        <a href="https://github.com/SpaceXLaunchBot/" rel="noreferrer" target="_blank">GitHub</a>
        {/* TODO: Make logging out like on mee6's website instead of having button always there. */}
        <UserInfo discordOAuthToken={discordOAuthToken} loggedIn={loggedIn} />
        <Login loggedIn={loggedIn} logOut={logOut} />
      </Grid>
    </Grid>
  );
}
