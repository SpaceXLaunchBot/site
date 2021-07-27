import React from 'react';
import { Grid } from '@material-ui/core';

export default function Guild(props) {
  const { name, icon, children } = props;
  return (
    <div className="guild">
      <Grid
        container
        direction="row"
        justifyContent="center"
        alignItems="center"
        className="guildHeader"
      >
        <img className="circleImg guildIcon" alt="guild icon" src={icon} />
        <h2 className="guildName">{name}</h2>
      </Grid>
      {children}
    </div>
  );
}
