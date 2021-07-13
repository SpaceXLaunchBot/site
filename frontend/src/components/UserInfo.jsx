import React, { useEffect, useState } from 'react';

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

  return (
    <div className="userInfo">
      <img className="userName" src={userData.avatarUrl} alt={'User\'s avatar'} />
      <p>{userData.userName}</p>
    </div>
  );
}
