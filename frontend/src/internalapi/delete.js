export default async function deleteChannel(body) {
  const res = await fetch('/api/channel', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  });
  return res.json();
}
