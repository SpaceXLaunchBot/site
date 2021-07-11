export default async function getMetrics() {
  const res = await fetch('/api/metrics', {
    method: 'GET',
  });
  return res.json();
}
