import React, { useEffect, useState } from 'react';
import { Button } from '@material-ui/core';

let oauthUrl = 'https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=https%3A%2F%2Fslb.simonj.dev%2F&response_type=token&scope=identify%20guilds';
if (window.location.hostname === 'localhost') {
  oauthUrl = `https://discord.com/api/oauth2/authorize?client_id=782810710546579476&redirect_uri=http%3A%2F%2Flocalhost%3A${window.location.host.split(':')[1]}&response_type=token&scope=identify%20guilds`;
}

async function getUserInfo(token) {
  const res = await fetch('https://discord.com/api/users/@me', {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  const json = await res.json();
  return {
    userName: json.username,
    avatarUrl: `https://cdn.discordapp.com/avatars/${json.id}/${json.avatar}.png`,
  };
}

export default function Login(props) {
  const { discordOAuthToken, loggedIn, logOut } = props;
  const [userData, setUserData] = useState({});

  useEffect(async () => {
    if (loggedIn) {
      const storedUD = JSON.parse(localStorage.getItem('user-data'));
      if (storedUD !== null) {
        console.log('Loading cached userData');
        setUserData(storedUD);
      } else {
        const newUserData = await getUserInfo(discordOAuthToken);
        setUserData(newUserData);
        localStorage.setItem('user-data', JSON.stringify(newUserData));
      }
    }
  }, [loggedIn]);

  // TODO: Swap returns
  if (loggedIn) {
    return (
      <div className="userInfo">
        <p>{userData.userName}</p>
        <img src={userData.avatarUrl} alt={'User\'s avatar'} />
        {/* <p className="logOutHoverer">▼</p> */}
        <Button
          variant="contained"
          color="secondary"
          className="logOutBtn"
          onClick={logOut}
        >
          Log Out
        </Button>
      </div>
    );
  }

  return (
  // eslint-disable-next-line no-return-assign
    <Button variant="contained" color="secondary" onClick={() => window.location.href = oauthUrl}>Login</Button>
  );
}
