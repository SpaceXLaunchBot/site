export default async function getStats() {
  const res = await fetch('/api/metrics', {
    method: 'GET',
  });
  return res.json();
}
