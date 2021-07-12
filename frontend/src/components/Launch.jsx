import React, { useEffect, useState } from 'react';
import moment from 'moment';
import Loader from 'react-loader-spinner';
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
    return (
      <Loader
        className="loader"
        type="Grid"
        color="#00BFFF"
        height={25}
        width={25}
      />
    );
  }

  const launchMoment = moment(launchInfo.date_utc);

  let img;
  if (launchInfo.links.patch.small !== null) {
    img = <img src={launchInfo.links.patch.small} className="launchPatch" alt={`${launchInfo.name} mission patch`} />;
  }

  return (
    <div>
      <h1>{launchInfo.name}</h1>
      {img}
      <p>{`Launching ${launchMoment.format('D MMM YYYY [at] HH:mm')}`}</p>
      <p>{launchInfo.details}</p>
    </div>
  );
}
