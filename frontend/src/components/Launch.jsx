import React, { useEffect, useState } from 'react';
import moment from 'moment';
import getNextLaunch from '../spacexapi/nextlaunch';
import '../css/Launch.scss';

export default function Launch() {
  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [launchInfo, setLaunchInfo] = useState({});

  useEffect(() => {
    getNextLaunch()
      .then((json) => {
        if (json === {}) {
          setError({ message: 'SpaceX API request failed' });
        } else {
          setLaunchInfo(json);
        }
        setIsLoaded(true);
      });
  }, []);

  if (error) {
    return (
      <div>
        {`Error: ${error.message}`}
      </div>
    );
  } if (!isLoaded) {
    return <div>Loading...</div>;
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
