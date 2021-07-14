import React from 'react';
import '../css/Home.scss';
import {
  IoRocketSharp, MdNotificationsActive, MdNotifications, AiTwotoneSetting,
} from 'react-icons/all';
import Invite from '../components/Invite';
import Feature from '../components/Feature';

export default function Home() {
  // TODO: Can we add the icon className inside Feature?
  return (
    <div className="home">
      <div className="welcome">
        <h2>News, information, and notifications about SpaceX launches</h2>
        <Invite />
      </div>
      <div className="features">
        <Feature
          icon={<IoRocketSharp className="featureIcon" />}
          feature="See Launch Information"
          description="See information about any previous or planned launch"
        />
        <Feature
          icon={<MdNotificationsActive className="featureIcon" />}
          feature="Launch Notifications"
          description="Subscribe a channel to the notification service and never miss a launch again!"
        />
        <Feature
          icon={<MdNotifications className="featureIcon" />}
          feature="Launch Changes"
          description="Get notified when an upcoming launch changes"
        />
        <Feature
          icon={<AiTwotoneSetting className="featureIcon" />}
          feature="Edit Settings"
          description="All your settings for the bot can be changed on this website, just log in and go to the 'Server Settings' tab"
        />
      </div>
    </div>
  );
}
