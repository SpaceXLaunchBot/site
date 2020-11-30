import React from 'react';
import moment from 'moment';

export default function Launch(props) {
    const { launchInfo } = props;

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
