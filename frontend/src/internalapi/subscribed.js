export default async function getSubscribed() {
  const res = await fetch('/api/subscribed', {
    method: 'GET',
  });
  return res.json();
}
