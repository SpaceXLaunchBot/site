import React from 'react';
import { Box, Grid } from '@material-ui/core';

export default function Guild(props) {
  const { name, icon, children } = props;
  return (
    <Box className="guild">
      <Grid
        container
        direction="row"
        justify="center"
        alignItems="center"
        className="guildHeader"
      >
        <img className="guildIcon" alt="guild icon" src={icon} />
        <h2 className="guildName">{name}</h2>
      </Grid>
      {children}
    </Box>
  );
}
