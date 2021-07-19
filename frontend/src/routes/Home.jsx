import React from 'react';
import '../css/Home.scss';
import { useLocation } from 'react-router-dom';
import Invite from '../components/Invite';
import Feature from '../components/Feature';

// TODO: h2 should take up same width as features.

export default function Home() {
  const params = new URLSearchParams(useLocation().search);
  let err = '';

  // TODO: Properly report error to user.
  if (params.has('login_error')) {
    err = params.get('login_error');
  }

  return (
    <div className="home">
      <div className="welcome">
        <h2>News, information, and notifications about SpaceX launches</h2>
        <Invite />
        {err !== '' && <p>{`Login failed with error: ${err}`}</p>}
      </div>
      <div className="features">
        <Feature
          icon="info"
          feature="See Launch Information"
          description="See information about any previous or planned launch"
        />
        <Feature
          icon="notifications_active"
          feature="Launch Notifications"
          description="Subscribe a channel to the notification service and never miss a launch again!"
        />
        <Feature
          icon="notifications"
          feature="Launch Changes"
          description="Get notified when an upcoming launch changes"
        />
        <Feature
          icon="settings"
          feature="Edit Settings"
          description="All your settings for the bot can be changed on this website, just log in and go to the 'Server Settings' tab"
        />
      </div>
    </div>
  );
}
