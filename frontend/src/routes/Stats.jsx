import React, { useEffect, useState } from 'react';
import {
  CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis,
} from 'recharts';
import useMediaQuery from '@material-ui/core/useMediaQuery';
import Loader from '../components/Loader';
import getStats from '../internalapi/stats';
import '../css/Stats.scss';

export default function Stats() {
  const [counts, setCounts] = useState([]);
  // const [actionCounts, setActionCounts] = useState([]);
  const [loaded, setLoaded] = useState(false);
  const [error, setError] = useState('');
  // Maybe this should be passed down thru props to reduce useMediaQuery usage?
  const lessThan750px = useMediaQuery('(max-width:750px)');

  useEffect(async () => {
    try {
      const s = await getStats();
      setCounts(s.counts);
      // setActionCounts(s.action_counts);
    } catch (e) {
      if (e.constructor === SyntaxError) {
        setError('Server returned invalid JSON');
      } else {
        setError(e.toString());
      }
    }
    setLoaded(true);
  }, []);

  if (!loaded) {
    return <Loader />;
  }

  if (error !== '') {
    return (
      <div>
        <h2>Failed to get data</h2>
        <p>{error}</p>
      </div>
    );
  }

  // TODO: Look into ResponsiveContainer.
  const chartWidths = lessThan750px ? 400 : 600;

  return (
    <div>
      <LineChart className="chart" width={chartWidths} height={250} data={counts}>
        <Legend verticalAlign="top" height={36} />
        <Line strokeWidth={2} dot={false} name="Server Count" type="monotone" dataKey="g" stroke="#a7a3ff" />
        <CartesianGrid stroke="#686D73" />
        <XAxis tickMargin={10} dataKey="d" />
        <YAxis tickCount={6} type="number" domain={['dataMin - 10', 'dataMax + 10']} />
        <Tooltip />
      </LineChart>
      <LineChart className="chart" width={chartWidths} height={250} data={counts}>
        <Legend verticalAlign="top" height={36} />
        <Line strokeWidth={2} dot={false} name="Subscribed Channel Count" type="monotone" dataKey="s" stroke="#13f088" />
        <CartesianGrid stroke="#686D73" />
        <XAxis tickMargin={10} dataKey="d" />
        <YAxis tickCount={6} type="number" domain={['dataMin - 10', 'dataMax + 10']} />
        <Tooltip />
      </LineChart>
    </div>
  );
}
