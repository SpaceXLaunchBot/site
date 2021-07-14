import React from 'react';
import '../css/Feature.scss';

export default function Feature(props) {
  const { icon, feature, description } = props;
  return (
    <div className="feature">
      <div className="featureHeader">
        {icon}
        <h3>{feature}</h3>
      </div>
      <p>{description}</p>
    </div>
  );
}
