export default async function deleteChannel(token, body) {
  const res = await fetch('/api/channel', {
    method: 'DELETE',
    headers: {
      Authorization: token,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  });
  return res.json();
}
