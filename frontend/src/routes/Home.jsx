import React from 'react';
import Launch from '../components/Launch';

export default function Home() {
  const inviteClicked = () => {
    console.log('clicked');
  };

  return (
    <div>
      <div className="invite" onClick={inviteClicked} onKeyDown={inviteClicked} role="button" tabIndex={0}>
        <img src="/discordlogo.svg" alt="Discord Icon" />
        <p>Add SpaceXLaunchBot</p>
      </div>
      <Launch />
    </div>
  );
}
