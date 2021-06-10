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
import updateChannel from '../internalapi/update';
import deleteChannel from '../internalapi/delete';

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
    marginLeft: '0',
    marginRight: '0.5em',
  },
}));

export default function Channel(props) {
  const { info, guildId, discordOAuthToken } = props;
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

  const saveBtnClicked = async () => {
    addToast(`Saving settings for ${info.name}`, { appearance: 'info' });
    const body = {
      id: info.id,
      guild_id: guildId,
      notification_type: notificationType,
      launch_mentions: launchMentions,
    };
    const json = await updateChannel(discordOAuthToken, body);
    if (json.success === true) {
      addToast(`Saved settings for ${info.name}`, { appearance: 'success' });
    } else {
      addToast(`Error saving settings for ${info.name}: ${json.error}`, { appearance: 'error' });
    }
  };

  const deleteBtnClicked = async () => {
    addToast(`Unsubscribing channel ${info.name}`, { appearance: 'info' });
    const body = {
      id: info.id,
      guild_id: guildId,
    };
    const json = await deleteChannel(discordOAuthToken, body);
    if (json.success === true) {
      addToast(`Unsubscribed ${info.name}`, { appearance: 'success' });
    } else {
      addToast(`Error unsubscribing channel ${info.name}: ${json.error}`, { appearance: 'error' });
    }
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
      <Grid item xs={12} md={6}>
        <FormControl component="fieldset">
          <RadioGroup row name="subscription-type" value={notificationType} onChange={radioChanged}>
            <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="all" control={<Radio />} label="All" />
            <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="schedule" control={<Radio />} label="Schedule" />
            <FormControlLabel classes={{ root: classes.radioButton }} labelPlacement="top" value="launch" control={<Radio />} label="Launch" />
          </RadioGroup>
        </FormControl>
      </Grid>
      <Grid item xs={12} md={6}>
        <TextField
          classes={{ root: classes.textField }}
          InputLabelProps={{ classes: { root: classes.inputPlaceholder } }}
          id="standard-basic"
          label="Launch Mentions"
          value={launchMentions}
          onChange={textFieldChanged}
        />
      </Grid>
      <Grid item xs={12}>
        <IconButton color="secondary" onClick={saveBtnClicked}>
          <Icon>save</Icon>
        </IconButton>
        <IconButton color="secondary" onClick={deleteBtnClicked}>
          <Icon>delete</Icon>
        </IconButton>
      </Grid>
    </Grid>
  );
}
