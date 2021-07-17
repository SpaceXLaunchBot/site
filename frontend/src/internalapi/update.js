export default async function updateChannel(body) {
  const res = await fetch('/api/channel', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  });
  return res.json();
}
