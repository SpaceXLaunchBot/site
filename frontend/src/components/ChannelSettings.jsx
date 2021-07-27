import React, { useState } from 'react';
import {
  FormControl,
  FormControlLabel,
  Icon,
  IconButton,
  Radio,
  RadioGroup,
  TextField,
} from '@material-ui/core';
import { useToasts } from 'react-toast-notifications';
import updateChannel from '../internalapi/update';
import deleteChannel from '../internalapi/delete';

// TODO: Delete should remove the rendered channel.

export default function ChannelSettings(props) {
  const {
    guildId, guildName, guildIcon, channelInfo,
  } = props;
  const [notificationType, setNotificationType] = useState(channelInfo.notification_type);
  const [launchMentions, setLaunchMentions] = useState(channelInfo.launch_mentions);
  const { addToast } = useToasts();

  const textFieldChanged = (e) => {
    setLaunchMentions(e.target.value);
  };
  const radioChanged = (e) => {
    setNotificationType(e.target.value);
  };

  const saveBtnClicked = async () => {
    addToast(`Saving settings for ${channelInfo.name}`, { appearance: 'info' });
    const body = {
      id: channelInfo.id,
      guild_id: guildId,
      notification_type: notificationType,
      launch_mentions: launchMentions,
    };
    const json = await updateChannel(body);
    if (json.success === true) {
      addToast(`Saved settings for ${channelInfo.name}`, { appearance: 'success' });
    } else {
      addToast(`Error saving settings for ${channelInfo.name}: ${json.error}`, { appearance: 'error' });
    }
  };

  const deleteBtnClicked = async () => {
    addToast(`Unsubscribing channel ${channelInfo.name}`, { appearance: 'info' });
    const body = {
      id: channelInfo.id,
      guild_id: guildId,
    };
    const json = await deleteChannel(body);
    if (json.success === true) {
      addToast(`Unsubscribed ${channelInfo.name}`, { appearance: 'success' });
    } else {
      addToast(`Error unsubscribing channel ${channelInfo.name}: ${json.error}`, { appearance: 'error' });
    }
  };

  return (
    <div className="channelSettings">
      <div className="channelSettingsGuild">
        <img className="circleImg guildIcon" alt="guild icon" src={guildIcon} />
        <h2 className="guildName">{guildName}</h2>
      </div>
      <div className="channelSettingsChannel">
        <h3>{`#${channelInfo.name}`}</h3>
      </div>
      <div className="channelSettingsSettings">
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
        <TextField
          label="Launch Mentions"
          value={launchMentions}
          onChange={textFieldChanged}
          className="launchMentionsInput"
          multiline
        />
      </div>
      <div className="channelSettingsButtons">
        <IconButton onClick={saveBtnClicked}>
          <Icon>save</Icon>
        </IconButton>
        <IconButton onClick={deleteBtnClicked}>
          <Icon>delete</Icon>
        </IconButton>
      </div>
    </div>
  );
}
