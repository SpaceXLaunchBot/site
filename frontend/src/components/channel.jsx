import React, { useState } from 'react';
import {
    FormControl,
    FormControlLabel,
    Grid,
    Icon,
    IconButton,
    Radio,
    RadioGroup,
    TextField,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import { useToasts } from 'react-toast-notifications';

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
    const { info, discordOAuthToken } = props;
    const [notificationType, setNotificationType] = useState(info.notification_type);
    const [launchMentions, setLaunchMentions] = useState(info.launch_mentions.String);
    const { addToast } = useToasts();
    const classes = useStyles();

    const textFieldChanged = (e) => {
        setLaunchMentions(e.target.value);
    };
    const radioChanged = (e) => {
        setNotificationType(e.target.value);
    };

    const saveBtnClicked = () => {
        // Send POST request with discordOAuthToken.
        addToast(`Saving settings for ${info.name}`, { appearance: 'info' });
        addToast(`Saved settings for ${info.name}`, { appearance: 'success' });
        addToast(`Cannot save settings for ${info.name}`, { appearance: 'error' });
    };
    const deleteBtnClicked = () => {
        // Send DELETE request with discordOAuthToken.
    };

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
                <RadioGroup row name="subscription-type" value={notificationType} onChange={radioChanged}>
                    <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="all" control={<Radio />} label="All" />
                    <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="schedule" control={<Radio />} label="Schedule" />
                    <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="launch" control={<Radio />} label="Launch" />
                </RadioGroup>
            </FormControl>
            <TextField
                classes={{ root: classes.textField }}
                InputLabelProps={{ classes: { root: classes.inputPlaceholder } }}
                id="standard-basic"
                label="Launch Mentions"
                value={launchMentions}
                onChange={textFieldChanged}
            />
            <IconButton color="secondary" onClick={saveBtnClicked}>
                <Icon>save</Icon>
            </IconButton>
            <IconButton color="secondary" onClick={deleteBtnClicked}>
                <Icon>delete</Icon>
            </IconButton>
        </Grid>
    );
}
