import React, { useEffect, useState } from 'react';
import '../css/UserInfo.scss';
import useMediaQuery from '@material-ui/core/useMediaQuery';

async function getUserData(token) {
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
export default function UserInfo(props) {
  const { discordOAuthToken, loggedIn } = props;
  const [userData, setUserData] = useState({});
  // TODO: Rename.
  const matches = useMediaQuery('(max-width:750px)');

  useEffect(async () => {
    if (loggedIn) {
      const storedUD = JSON.parse(localStorage.getItem('user-data'));
      if (storedUD !== null) {
        console.log('Using cached userData');
        setUserData(storedUD);
      } else {
        const newUserData = await getUserData(discordOAuthToken);
        setUserData(newUserData);
        localStorage.setItem('user-data', JSON.stringify(newUserData));
      }
    }
  }, [loggedIn]);

  if (!loggedIn) {
    return <div className="invisible" />;
  }

  const classes = matches ? 'circleImg avatarImgSmall' : 'circleImg';

  return (
    <div className="userInfo">
      {/* TODO: Make this smaller when matches is true */}
      <img className={classes} src={userData.avatarUrl} alt={'User\'s avatar'} />
      {/* https://reactjs.org/docs/conditional-rendering.html#inline-if-with-logical--operator */}
      {matches === false && <p className="userName">{userData.userName}</p>}
    </div>
  );
}
