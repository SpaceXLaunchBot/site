import React from 'react';
import '../css/Invite.scss';

const inviteUrl = 'https://discord.com/oauth2/authorize?client_id=411618411169447950&scope=bot&permissions=19456';

export default function Invite() {
  const inviteClicked = () => {
    window.open(inviteUrl, '_blank');
  };

  return (
    <div className="invite" onClick={inviteClicked} onKeyDown={inviteClicked} role="button" tabIndex={0}>
      <img src="/discordlogo.svg" alt="Discord Icon" />
      <p>Add SpaceXLaunchBot</p>
    </div>
  );
}
