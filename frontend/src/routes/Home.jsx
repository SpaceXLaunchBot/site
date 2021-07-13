import React from 'react';
import '../css/Home.scss';
import Invite from '../components/Invite';

export default function Home() {
  return (
    <div>
      <h2>News, information, and notifications about SpaceX launches</h2>
      <Invite />
    </div>
  );
}
