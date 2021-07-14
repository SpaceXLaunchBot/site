import React, { useEffect, useState } from 'react';
import {
  CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis,
} from 'recharts';
import Loader from 'react-loader-spinner';
import getStats from '../internalapi/stats';
import '../css/Stats.scss';

export default function Stats() {
  const [counts, setCounts] = useState({});
  const [loaded, setLoaded] = useState(false);
  const [error, setError] = useState('');

  useEffect(async () => {
    try {
      const s = await getStats();
      setCounts(s.counts);
    } catch (e) {
      setError(e.toString());
    }
    setLoaded(true);
  }, []);

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

  if (error !== '') {
    return (
      <div>
        <h2>Failed to get data</h2>
        <p>{error}</p>
      </div>
    );
  }

  return (
    <LineChart className="chart" width={600} height={300} data={counts}>
      <Legend verticalAlign="top" height={36} />
      <Line name="Server Count" type="monotone" dataKey="g" stroke="#a7a3ff" />
      <Line name="Subscribed Channel Count" type="monotone" dataKey="s" stroke="#13f088" />
      <CartesianGrid stroke="#ccc" />
      <XAxis dataKey="d" />
      <YAxis type="number" domain={['dataMin - 100', 'dataMax + 100']} />
      <Tooltip />
    </LineChart>
  );
}
