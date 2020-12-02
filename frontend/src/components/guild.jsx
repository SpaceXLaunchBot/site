import React from 'react';
import { Box, Grid } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles(() => ({
    box: {
        borderStyle: 'solid',
        borderRadius: '1em',
        borderColor: '#393e43',
        marginTop: '0.5em;',
        marginBottom: '0.5em;',
    },
    guildHeader: {
        marginTop: '0.5em',
        marginBottom: '0.5em',
    },
}));

export default function Guild(props) {
    const { name, icon, children } = props;
    const classes = useStyles();
    return (
        <Box classes={{ root: classes.box }}>
            <Grid
                container
                direction="row"
                justify="center"
                alignItems="center"
                classes={{ root: classes.guildHeader }}
            >
                <img className="guildIcon" alt="" src={icon} />
                <h2>{name}</h2>
            </Grid>
            {children}
        </Box>
    );
}
