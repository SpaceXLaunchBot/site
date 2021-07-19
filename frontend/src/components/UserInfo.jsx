import React, { useEffect, useState } from 'react';
import '../css/UserInfo.scss';
import useMediaQuery from '@material-ui/core/useMediaQuery';

export default function UserInfo(props) {
  const { loggedIn } = props;
  const [userData, setUserData] = useState({});
  const lessThan820px = useMediaQuery('(max-width:820px)');

  useEffect(async () => {
    if (loggedIn) {
      const storedUD = JSON.parse(localStorage.getItem('user-data'));
      if (storedUD !== null) {
        console.log('Using cached userData');
        setUserData(storedUD);
      } else {
        const res = await fetch('/api/userinfo');
        const json = await res.json();
        if (json.success === true) {
          const newUserData = json.user_info;
          setUserData(newUserData);
          localStorage.setItem('user-data', JSON.stringify(newUserData));
        } else {
          // TODO: Give a toast popup or something.
        }
      }
    }
  }, [loggedIn]);

  if (!loggedIn) {
    return <div className="invisible" />;
  }

  const classes = lessThan820px ? 'circleImg userAvatarSmall' : 'circleImg';

  return (
    <div className="userInfo">
      <img className={classes} src={userData.avatar_url} alt={'User\'s avatar'} />
      {/* https://reactjs.org/docs/conditional-rendering.html#inline-if-with-logical--operator */}
      {lessThan820px === false && <p className="userName">{userData.username}</p>}
    </div>
  );
}
