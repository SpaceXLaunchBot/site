import React, { useState } from 'react';
import {
    Grid, FormControl, IconButton, TextField, Radio, RadioGroup, FormControlLabel, Button,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import Icon from '@material-ui/core/Icon';

const useStyles = makeStyles((theme) => ({
    grid: {
        color: theme.palette.text.primary,
        backgroundColor: '#393e43',
        margin: '0 auto', // centers it
        marginBottom: '1em', // space at bottom
        borderRadius: '1em',
        width: '75%',
    },
    textField: {
        marginBottom: '1em', // space at bottom when screen width is small
        marginLeft: '0.5em',
        marginRight: '0.5em',
        '& label.Mui-focused': {
            color: theme.palette.text.primary,
        },
    },
    inputPlaceholder: {
        color: theme.palette.text.secondary,
    },
    radioButton: {
        margin: '0.5em',
    },
}));

export default function Channel(props) {
    const { info } = props;
    const classes = useStyles();
    return (
        <Grid
            container
            direction="row"
            justify="center"
            alignItems="center"
            classes={{ root: classes.grid }}
        >
            <Grid item xs={12}>
                <h3>{info.name}</h3>
            </Grid>
            <FormControl component="fieldset">
                <RadioGroup row name="subscription-type" defaultValue={info.notification_type}>
                    <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="all" control={<Radio />} label="All" />
                    <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="schedule" control={<Radio />} label="Schedule" />
                    <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="launch" control={<Radio />} label="Launch" />
                </RadioGroup>
            </FormControl>
            <TextField
                classes={{ root: classes.textField }}
                InputLabelProps={{
                    classes: { root: classes.inputPlaceholder },
                }}
                id="standard-basic"
                label="Launch Mentions"
                defaultValue={info.launch_mentions.String}
            />
            <IconButton
                color="secondary"
            >
                <Icon>save</Icon>
            </IconButton>
            <IconButton
                color="secondary"
            >
                <Icon>delete</Icon>
            </IconButton>
        </Grid>
    );
}
