export default function GetGuildsWithSubscribed(token) {
  return fetch('/api/subscribed', {
    method: 'GET',
    headers: {
      authorization: token,
    },
  })
    .then((res) => res.json())
    .then(
      (json) => json,
      () => ({}),
    );
}
