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
// import useMediaQuery from '@material-ui/core/useMediaQuery';
import Loader from '../components/Loader';
// import getStats from '../internalapi/stats';
import '../css/Stats.scss';

const testData = '{"success":true,"status_code":200,"counts":[{"g":662,"s":355,"d":"2021-07-10 00:00:00"},{"g":662,"s":355,"d":"2021-07-10 12:00:00"},{"g":662,"s":355,"d":"2021-07-11 00:00:00"},{"g":662,"s":355,"d":"2021-07-11 12:00:00"},{"g":662,"s":355,"d":"2021-07-12 00:00:00"},{"g":663,"s":355,"d":"2021-07-12 12:00:00"},{"g":662,"s":355,"d":"2021-07-13 00:00:00"},{"g":662,"s":355,"d":"2021-07-13 01:00:00"},{"g":662,"s":355,"d":"2021-07-13 02:00:00"},{"g":662,"s":355,"d":"2021-07-13 03:00:00"},{"g":662,"s":355,"d":"2021-07-13 04:00:00"},{"g":662,"s":355,"d":"2021-07-13 05:00:00"},{"g":662,"s":355,"d":"2021-07-13 06:00:00"},{"g":662,"s":355,"d":"2021-07-13 07:00:00"},{"g":662,"s":355,"d":"2021-07-13 08:00:00"},{"g":662,"s":355,"d":"2021-07-13 09:00:00"},{"g":662,"s":355,"d":"2021-07-13 10:00:00"},{"g":662,"s":355,"d":"2021-07-13 11:00:00"},{"g":662,"s":356,"d":"2021-07-13 12:00:00"},{"g":662,"s":356,"d":"2021-07-13 13:00:00"},{"g":662,"s":356,"d":"2021-07-13 14:00:00"},{"g":663,"s":356,"d":"2021-07-13 15:00:00"},{"g":663,"s":356,"d":"2021-07-13 16:00:00"},{"g":663,"s":356,"d":"2021-07-13 17:00:00"},{"g":663,"s":356,"d":"2021-07-13 18:00:00"},{"g":663,"s":356,"d":"2021-07-13 19:00:00"},{"g":663,"s":356,"d":"2021-07-13 20:00:00"},{"g":663,"s":356,"d":"2021-07-13 21:00:00"},{"g":663,"s":356,"d":"2021-07-13 22:00:00"},{"g":664,"s":356,"d":"2021-07-13 23:00:00"},{"g":664,"s":356,"d":"2021-07-14 00:00:00"},{"g":664,"s":356,"d":"2021-07-14 01:00:00"},{"g":664,"s":356,"d":"2021-07-14 02:00:00"},{"g":664,"s":356,"d":"2021-07-14 03:00:00"},{"g":664,"s":356,"d":"2021-07-14 04:00:00"},{"g":664,"s":356,"d":"2021-07-14 05:00:00"},{"g":665,"s":356,"d":"2021-07-14 06:00:00"},{"g":664,"s":356,"d":"2021-07-14 07:00:00"},{"g":664,"s":356,"d":"2021-07-14 08:00:00"},{"g":664,"s":356,"d":"2021-07-14 09:00:00"},{"g":664,"s":356,"d":"2021-07-14 10:00:00"},{"g":664,"s":356,"d":"2021-07-14 11:00:00"},{"g":664,"s":356,"d":"2021-07-14 12:00:00"},{"g":664,"s":356,"d":"2021-07-14 13:00:00"},{"g":664,"s":356,"d":"2021-07-14 14:00:00"},{"g":664,"s":356,"d":"2021-07-14 15:00:00"},{"g":664,"s":356,"d":"2021-07-14 16:00:00"},{"g":664,"s":356,"d":"2021-07-14 17:00:00"},{"g":664,"s":356,"d":"2021-07-14 18:00:00"},{"g":664,"s":356,"d":"2021-07-14 19:00:00"},{"g":664,"s":354,"d":"2021-07-14 20:00:00"},{"g":664,"s":354,"d":"2021-07-14 21:00:00"},{"g":664,"s":354,"d":"2021-07-14 22:00:00"},{"g":664,"s":355,"d":"2021-07-14 23:00:00"},{"g":664,"s":355,"d":"2021-07-15 00:00:00"},{"g":664,"s":355,"d":"2021-07-15 01:00:00"},{"g":664,"s":355,"d":"2021-07-15 02:00:00"},{"g":664,"s":355,"d":"2021-07-15 03:00:00"},{"g":664,"s":355,"d":"2021-07-15 04:00:00"},{"g":664,"s":355,"d":"2021-07-15 05:00:00"},{"g":664,"s":355,"d":"2021-07-15 06:00:00"},{"g":664,"s":355,"d":"2021-07-15 07:00:00"},{"g":664,"s":355,"d":"2021-07-15 08:00:00"},{"g":664,"s":355,"d":"2021-07-15 09:00:00"},{"g":664,"s":355,"d":"2021-07-15 10:00:00"},{"g":664,"s":355,"d":"2021-07-15 11:00:00"},{"g":664,"s":355,"d":"2021-07-15 12:00:00"},{"g":665,"s":356,"d":"2021-07-15 13:00:00"},{"g":665,"s":356,"d":"2021-07-15 14:00:00"},{"g":665,"s":356,"d":"2021-07-15 15:00:00"},{"g":665,"s":357,"d":"2021-07-15 16:00:00"},{"g":665,"s":357,"d":"2021-07-15 17:00:00"},{"g":665,"s":357,"d":"2021-07-15 18:00:00"},{"g":665,"s":357,"d":"2021-07-15 19:00:00"},{"g":665,"s":357,"d":"2021-07-15 20:00:00"},{"g":664,"s":357,"d":"2021-07-15 21:00:00"},{"g":664,"s":357,"d":"2021-07-15 22:00:00"},{"g":664,"s":357,"d":"2021-07-15 23:00:00"},{"g":664,"s":357,"d":"2021-07-16 00:00:00"},{"g":663,"s":357,"d":"2021-07-16 01:00:00"},{"g":663,"s":357,"d":"2021-07-16 02:00:00"},{"g":663,"s":357,"d":"2021-07-16 03:00:00"},{"g":663,"s":357,"d":"2021-07-16 04:00:00"},{"g":663,"s":357,"d":"2021-07-16 05:00:00"},{"g":663,"s":357,"d":"2021-07-16 06:00:00"},{"g":664,"s":357,"d":"2021-07-16 07:00:00"},{"g":664,"s":357,"d":"2021-07-16 08:00:00"},{"g":664,"s":357,"d":"2021-07-16 09:00:00"},{"g":664,"s":357,"d":"2021-07-16 10:00:00"},{"g":664,"s":357,"d":"2021-07-16 11:00:00"},{"g":664,"s":357,"d":"2021-07-16 12:00:00"},{"g":664,"s":357,"d":"2021-07-16 13:00:00"},{"g":664,"s":357,"d":"2021-07-16 14:00:00"},{"g":664,"s":357,"d":"2021-07-16 15:00:00"},{"g":664,"s":357,"d":"2021-07-16 16:00:00"},{"g":664,"s":357,"d":"2021-07-16 17:00:00"},{"g":664,"s":357,"d":"2021-07-16 18:00:00"},{"g":664,"s":357,"d":"2021-07-16 19:00:00"},{"g":664,"s":357,"d":"2021-07-16 20:00:00"},{"g":664,"s":357,"d":"2021-07-16 21:00:00"}],"action_counts":[{"a":"launch","c":4},{"a":"info","c":5},{"a":"guild_join","c":10},{"a":"guild_remove","c":13},{"a":"add","c":11},{"a":"help","c":8},{"a":"nextlaunch","c":19}]}';
const s = JSON.parse(testData);

export default function Stats() {
  const [counts, setCounts] = useState([]);
  const [actionCounts, setActionCounts] = useState([]);
  const [loaded, setLoaded] = useState(false);
  const [error, setError] = useState('');
  const [chartWidth, setChartWidth] = useState(0);

  useEffect(async () => {
    try {
      // const s = await getStats();
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
