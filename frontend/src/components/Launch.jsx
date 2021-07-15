import React, { useEffect, useState } from 'react';
import moment from 'moment';
import Loader from './Loader';
import getNextLaunch from '../spacexapi/nextlaunch';
import '../css/Launch.scss';

export default function Launch() {
  const [error, setError] = useState('');
  const [loaded, setLoaded] = useState(false);
  const [launchInfo, setLaunchInfo] = useState({});

  useEffect(async () => {
    const json = await getNextLaunch();
    if (json === {}) {
      setError('SpaceX API request failed');
    } else {
      setLaunchInfo(json);
    }
    setLoaded(true);
  }, []);

  if (error !== '') {
    return (
      <div>
        {`Error: ${error}`}
      </div>
    );
  }

  if (!loaded) {
    return <Loader />;
  }

  const launchMoment = moment(launchInfo.date_utc);

  let img;
  if (launchInfo.links.patch.small !== null) {
    img = <img src={launchInfo.links.patch.small} className="launchPatch" alt={`${launchInfo.name} mission patch`} />;
  }

  return (
    <div className="launch">
      <h2>Next Launch</h2>
      <p>{launchInfo.name}</p>
      {img}
      <p>{`Launching ${launchMoment.format('D MMM YYYY [at] HH:mm')}`}</p>
      <p>{launchInfo.details}</p>
    </div>
  );
}
