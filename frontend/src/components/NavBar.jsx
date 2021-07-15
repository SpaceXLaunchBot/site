import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { Drawer, Grid, MenuItem } from '@material-ui/core';
import { AiOutlineMenu } from 'react-icons/all';
import useMediaQuery from '@material-ui/core/useMediaQuery';
import Login from './Login';
import '../css/NavBar.scss';
import UserInfo from './UserInfo';

export default function NavBar(props) {
  const { discordOAuthToken, loggedIn, logOut } = props;
  const [drawerOpen, setDrawerOpen] = useState(false);
  const lessThan600px = useMediaQuery('(max-width:600px)');

  const drawerOpenBtnClicked = () => {
    setDrawerOpen(true);
  };

  const closeDrawer = () => {
    setDrawerOpen(false);
  };

  // TODO:
  // The Grid onClick works but it closes the drawer wherever you click. It might be nicer to not
  // close it when you select a menu item, since you can see them in the background a user might be
  // clicking through multiple pages to see what they are.
  // ALSO: Not sure about 2 copies of all links for header and drawer, maybe some way to DRY?

  return (
    <Grid
      container
      direction="row"
      justify="space-evenly"
      alignItems="stretch"
      className="navbar"
      onClick={drawerOpen ? closeDrawer : null}
    >
      <Grid item xs={12} sm={3} className="navbarTitle">
        <img className="circleImg" alt="SpaceXLaunchBot icon" src="/logo192.png" />
        <Link to="/"><h1>SpaceXLaunchBot</h1></Link>
        {lessThan600px && <AiOutlineMenu className="drawerOpener" onClick={drawerOpenBtnClicked} />}
      </Grid>
      {!lessThan600px
      && (
      <Grid item xs={12} sm={9} className="navbarLinks">
        <Link to="/commands">Commands</Link>
        <Link to="/settings">Server Settings</Link>
        <Link to="/stats">Stats</Link>
        <a href="https://top.gg/bot/411618411169447950/" rel="noreferrer" target="_blank">Top.gg</a>
        <a href="https://github.com/SpaceXLaunchBot/" rel="noreferrer" target="_blank">GitHub</a>
        <UserInfo discordOAuthToken={discordOAuthToken} loggedIn={loggedIn} />
        <Login loggedIn={loggedIn} logOut={logOut} />
      </Grid>
      )}
      {lessThan600px
      && (
        <Drawer className="menuDrawer" width={400} anchor="right" open={drawerOpen}>
          <Link to="/"><MenuItem>Home</MenuItem></Link>
          <Link to="/commands"><MenuItem>Commands</MenuItem></Link>
          <Link to="/settings"><MenuItem>Server Settings</MenuItem></Link>
          <Link to="/stats"><MenuItem>Stats</MenuItem></Link>
          <a
            href="https://top.gg/bot/411618411169447950/"
            rel="noreferrer"
            target="_blank"
          >
            <MenuItem>
              Top.gg
            </MenuItem>
          </a>
          <a
            href="https://github.com/SpaceXLaunchBot/"
            rel="noreferrer"
            target="_blank"
          >
            <MenuItem>
              GitHub
            </MenuItem>
          </a>
          <MenuItem>
            <UserInfo
              discordOAuthToken={discordOAuthToken}
              loggedIn={loggedIn}
            />
            <Login loggedIn={loggedIn} logOut={logOut} />
          </MenuItem>
        </Drawer>
      )}
    </Grid>
  );
}
