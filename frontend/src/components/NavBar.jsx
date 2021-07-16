import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import {
  Drawer, Grid, Icon, MenuItem,
} from '@material-ui/core';
import useMediaQuery from '@material-ui/core/useMediaQuery';
import Login from './Login';
import '../css/NavBar.scss';
import UserInfo from './UserInfo';

function ExternalLink(props) {
  const { href, text, children } = props;
  // We accept children so a MenuItem can be inserted.
  return (
    <a
      href={href}
      rel="noreferrer"
      target="_blank"
    >
      {text}
      {children}
    </a>
  );
}

export default function NavBar(props) {
  const { discordOAuthToken, loggedIn, logOut } = props;
  const [drawerOpen, setDrawerOpen] = useState(false);
  const lessThan600px = useMediaQuery('(max-width:600px)');
  const lessThan750px = useMediaQuery('(max-width:750px)');

  const drawerOpenClicked = () => {
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
  // ALSO: Logout menuitem isn't nested inside so user has to click on svg which is bad maybe

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
        <img className="botAvatar circleImg" alt="SpaceXLaunchBot icon" src="/logo192.png" />
        <Link to="/"><h1>SpaceXLaunchBot</h1></Link>
        {lessThan600px && <Icon className="drawerOpener" onClick={drawerOpenClicked}>menu</Icon>}
      </Grid>
      {!lessThan600px
      && (
      <Grid item xs={12} sm={9} className="navbarLinks">
        <Link to="/commands">Commands</Link>
        <Link to="/settings">Server Settings</Link>
        <Link to="/stats">Stats</Link>
        {!lessThan750px
        && ([
          <ExternalLink href="https://top.gg/bot/411618411169447950/" text="Top.gg" key={0} />,
          <ExternalLink href="https://github.com/SpaceXLaunchBot/" text="GitHub" key={1} />,
          <ExternalLink href="https://www.buymeacoffee.com/psidex" text="Donate" key={2} />,
        ])}
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
        <ExternalLink href="https://top.gg/bot/411618411169447950/">
          <MenuItem>Top.gg</MenuItem>
        </ExternalLink>
        <ExternalLink href="https://github.com/SpaceXLaunchBot/">
          <MenuItem>GitHub</MenuItem>
        </ExternalLink>
        <ExternalLink href="https://www.buymeacoffee.com/psidex">
          <MenuItem>Donate</MenuItem>
        </ExternalLink>
        <MenuItem>
          <UserInfo discordOAuthToken={discordOAuthToken} loggedIn={loggedIn} />
          <Login loggedIn={loggedIn} logOut={logOut} />
        </MenuItem>
      </Drawer>
      )}
    </Grid>
  );
}
