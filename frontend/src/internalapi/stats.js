export default async function getStats() {
  const res = await fetch('/api/stats', {
    method: 'GET',
  });
  return res.json();
}
