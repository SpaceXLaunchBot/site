import React from 'react';
import { Box, Grid } from '@material-ui/core';

export default function Guild(props) {
  const { name, icon, children } = props;
  return (
    <Box>
      <Grid
        container
        direction="row"
        justify="center"
        alignItems="center"
      >
        <img className="guildIcon" alt="" src={icon} />
        <h2>{name}</h2>
      </Grid>
      {children}
    </Box>
  );
}
