export default function UpdateSubscribedChannel(token, body) {
  return fetch('http://localhost:8080/api/updatesubscribedchannel', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Discord-Bearer-Token': token,
    },
    body: JSON.stringify(body),
  })
    .then((res) => res.json())
    .then(
      (json) => json,
      () => ({}),
    );
}
