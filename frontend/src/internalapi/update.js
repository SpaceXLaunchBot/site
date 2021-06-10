export default async function updateChannel(token, body) {
  const res = fetch('/api/channel', {
    method: 'PUT',
    headers: {
      Authorization: token,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  });
  return res.json();
}
