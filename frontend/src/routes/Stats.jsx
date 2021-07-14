import React, { useEffect, useState } from 'react';
import {
  CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis,
} from 'recharts';
import Loader from 'react-loader-spinner';
// import getStats from '../internalapi/stats';
import '../css/Stats.scss';
import useMediaQuery from '@material-ui/core/useMediaQuery';

const exampleData = '{"success":true,"status_code":200,"counts":[{"g":662,"s":355,"d":"2021-07-13 12:00:00"},{"g":662,"s":355,"d":"2021-07-13 01:00:00"},{"g":662,"s":355,"d":"2021-07-13 02:00:00"},{"g":662,"s":355,"d":"2021-07-13 03:00:00"},{"g":662,"s":355,"d":"2021-07-13 04:00:00"},{"g":662,"s":355,"d":"2021-07-13 05:00:00"},{"g":662,"s":355,"d":"2021-07-13 06:00:00"},{"g":662,"s":355,"d":"2021-07-13 07:00:00"},{"g":662,"s":355,"d":"2021-07-13 08:00:00"},{"g":662,"s":355,"d":"2021-07-13 09:00:00"},{"g":662,"s":355,"d":"2021-07-13 10:00:00"},{"g":662,"s":355,"d":"2021-07-13 11:00:00"},{"g":662,"s":356,"d":"2021-07-13 12:00:00"},{"g":662,"s":356,"d":"2021-07-13 01:00:00"},{"g":662,"s":356,"d":"2021-07-13 02:00:00"},{"g":663,"s":356,"d":"2021-07-13 03:00:00"},{"g":663,"s":356,"d":"2021-07-13 04:00:00"},{"g":663,"s":356,"d":"2021-07-13 05:00:00"},{"g":663,"s":356,"d":"2021-07-13 06:00:00"},{"g":663,"s":356,"d":"2021-07-13 07:00:00"},{"g":663,"s":356,"d":"2021-07-13 08:00:00"},{"g":663,"s":356,"d":"2021-07-13 09:00:00"},{"g":663,"s":356,"d":"2021-07-13 10:00:00"},{"g":664,"s":356,"d":"2021-07-13 11:00:00"},{"g":664,"s":356,"d":"2021-07-14 12:00:00"},{"g":664,"s":356,"d":"2021-07-14 01:00:00"},{"g":664,"s":356,"d":"2021-07-14 02:00:00"},{"g":664,"s":356,"d":"2021-07-14 03:00:00"},{"g":664,"s":356,"d":"2021-07-14 04:00:00"},{"g":664,"s":356,"d":"2021-07-14 05:00:00"},{"g":665,"s":356,"d":"2021-07-14 06:00:00"},{"g":664,"s":356,"d":"2021-07-14 07:00:00"},{"g":664,"s":356,"d":"2021-07-14 08:00:00"},{"g":664,"s":356,"d":"2021-07-14 09:00:00"},{"g":664,"s":356,"d":"2021-07-14 10:00:00"},{"g":664,"s":356,"d":"2021-07-14 11:00:00"},{"g":664,"s":356,"d":"2021-07-14 12:00:00"},{"g":664,"s":356,"d":"2021-07-14 01:00:00"},{"g":664,"s":356,"d":"2021-07-14 02:00:00"},{"g":664,"s":356,"d":"2021-07-14 03:00:00"},{"g":664,"s":356,"d":"2021-07-14 04:00:00"},{"g":664,"s":356,"d":"2021-07-14 05:00:00"}]}';
const exampleDataJson = JSON.parse(exampleData);

export default function Stats() {
  const [counts, setCounts] = useState({});
  const [loaded, setLoaded] = useState(false);
  const [error, setError] = useState('');
  // Maybe this should be passed down thru props to reduce useMediaQuery usage?
  const lessThan750px = useMediaQuery('(max-width:750px)');

  useEffect(async () => {
    try {
      // const s = await getStats();
      const s = exampleDataJson;
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
