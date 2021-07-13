import React from 'react';
import '../css/Invite.scss';

export default function Invite() {
  const inviteClicked = () => {
    console.log('clicked');
  };

  return (
    <div className="invite" onClick={inviteClicked} onKeyDown={inviteClicked} role="button" tabIndex={0}>
      <img src="/discordlogo.svg" alt="Discord Icon" />
      <p>Add SpaceXLaunchBot</p>
    </div>
  );
}
