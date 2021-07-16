import React, { useEffect, useState } from 'react';
import {
  CartesianGrid,
  Legend,
  Line,
  LineChart,
  PolarAngleAxis,
  PolarGrid,
  PolarRadiusAxis,
  Radar,
  RadarChart,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';
import Loader from '../components/Loader';
import getStats from '../internalapi/stats';
import '../css/Stats.scss';

export default function Stats() {
  const [counts, setCounts] = useState([]);
  const [actionCounts, setActionCounts] = useState([]);
  const [chartWidth, setChartWidth] = useState(0);
  const [loaded, setLoaded] = useState(false);
  const [error, setError] = useState('');

  useEffect(async () => {
    try {
      const s = await getStats();
      setCounts(s.counts);
      setActionCounts(s.action_counts);
    } catch (e) {
      if (e.constructor === SyntaxError) {
        setError('Server returned invalid JSON');
      } else {
        setError(e.toString());
      }
    }
    setLoaded(true);
  }, []);

  useEffect(() => {
    function handleResize() {
      if (window.innerWidth >= 750) {
        setChartWidth(600);
      } else if (window.innerWidth >= 500) {
        setChartWidth(window.innerWidth - 100);
      } else {
        setChartWidth(window.innerWidth);
      }
    }
    handleResize();
    // Idea from https://www.pluralsight.com/guides/re-render-react-component-on-window-resize
    window.addEventListener('resize', handleResize);
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

  return (
    <div className="stats">
      <h3>Server Count</h3>
      <LineChart className="chart" width={chartWidth} height={250} data={counts}>
        <Line strokeWidth={2} dot={false} name="Server Count" type="monotone" dataKey="g" stroke="#a7a3ff" />
        <CartesianGrid stroke="#686D73" />
        <XAxis tickMargin={10} dataKey="d" />
        <YAxis tickCount={6} type="number" width={35} domain={['dataMin - 10', 'dataMax + 10']} />
        <Tooltip />
      </LineChart>
      <h3>Subscribed Channel Count</h3>
      <LineChart className="chart" width={chartWidth} height={250} data={counts}>
        <Line strokeWidth={2} dot={false} name="Subscribed Channel Count" type="monotone" dataKey="s" stroke="#13f088" />
        <CartesianGrid stroke="#686D73" />
        <XAxis tickMargin={10} dataKey="d" />
        <YAxis tickCount={6} type="number" width={35} domain={['dataMin - 10', 'dataMax + 10']} />
        <Tooltip />
      </LineChart>
      <h3>Command Usage</h3>
      <RadarChart className="chart" outerRadius={100} width={chartWidth} height={250} data={actionCounts}>
        <Legend verticalAlign="bottom" />
        <PolarGrid stroke="#686D73" />
        <PolarAngleAxis dataKey="a" />
        <PolarRadiusAxis angle={30} domain={[0, Math.max(...Object.values(actionCounts))]} />
        <Radar name="Count" dataKey="c" stroke="#EB459E" fill="#ed2b93" fillOpacity={0.35} />
        <Tooltip />
      </RadarChart>
    </div>
  );
}
