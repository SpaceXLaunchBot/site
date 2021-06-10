export default async function getSubscribed(token) {
  const res = await fetch('/api/subscribed', {
    method: 'GET',
    headers: {
      Authorization: token,
    },
  });
  return res.json();
}
