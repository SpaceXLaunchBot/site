import React from 'react';
import { Link } from 'react-router-dom';
import Login from './Login';

export default function NavBar() {
  return (
    <div className="navbar">
      <div className="navbarTitle">
        <img alt="SpaceXLaunchBot icon" src="/logo192.png" />
        <Link to="/"><h1>SpaceXLaunchBot</h1></Link>
      </div>
      <div className="navbarLinks">
        <Link to="/settings">Bot Settings</Link>
        <Link to="/stats">Stats</Link>
        <Login />
      </div>
    </div>
  );
}
