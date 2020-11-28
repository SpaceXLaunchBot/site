import React from 'react';
import moment from 'moment';
import Countdown from './Countdown';

export default function Launch(props) {
    const { launchInfo } = props;

    const alreadyLaunched = launchInfo.date_unix <= Math.floor(Date.now() / 1000);
    const launchMoment = moment(launchInfo.date_utc);

    let date;
    if (alreadyLaunched) {
        date = <p>{launchMoment.format('MMM. D, YYYY [at] h:mm A z')}</p>;
    } else {
        date = <Countdown futureMoment={launchMoment} />;
    }

    let img;
    if (launchInfo.links.patch.small !== null) {
        img = <img src={launchInfo.links.patch.small} className="launchPatch" alt={`${launchInfo.name} mission patch`} />;
    }

    return (
        <div>
            <h1>{launchInfo.name}</h1>
            {img}
            <div className="inline">
                <p>{alreadyLaunched ? 'Launched on ' : 'Launching in '}</p>
                {date}
            </div>
            <p>{launchInfo.details}</p>
        </div>
    );
}
