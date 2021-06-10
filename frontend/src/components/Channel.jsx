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
import { useToasts } from 'react-toast-notifications';
import updateChannel from '../internalapi/update';
import deleteChannel from '../internalapi/delete';

export default function Channel(props) {
  const { info, guildId, discordOAuthToken } = props;
  const [notificationType, setNotificationType] = useState(info.notification_type);
  const [launchMentions, setLaunchMentions] = useState(info.launch_mentions);
  const { addToast } = useToasts();

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
      className="channelGrid"
    >
      <Grid item xs={12}>
        <h3>{info.name}</h3>
      </Grid>
      <Grid item xs={12} md={6}>
        <FormControl component="fieldset">
          <RadioGroup row name="subscription-type" value={notificationType} onChange={radioChanged}>
            <FormControlLabel
              labelPlacement="top"
              value="all"
              control={<Radio />}
              label="All"
              className="radioButton"
            />
            <FormControlLabel
              labelPlacement="top"
              value="schedule"
              control={<Radio />}
              label="Schedule"
              className="radioButton"
            />
            <FormControlLabel
              labelPlacement="top"
              value="launch"
              control={<Radio />}
              label="Launch"
              className="radioButton"
            />
          </RadioGroup>
        </FormControl>
      </Grid>
      <Grid item xs={12} md={6}>
        <TextField
          id="standard-basic"
          label="Launch Mentions"
          value={launchMentions}
          onChange={textFieldChanged}
          className="launchMentionsInput"
        />
      </Grid>
      <Grid item xs={12}>
        <IconButton onClick={saveBtnClicked}>
          <Icon>save</Icon>
        </IconButton>
        <IconButton onClick={deleteBtnClicked}>
          <Icon>delete</Icon>
        </IconButton>
      </Grid>
    </Grid>
  );
}
