export default function GetGuildsWithSubscribed(token) {
    return fetch('/api/guildswithsubscribed', {
        method: 'GET',
        headers: {
            'Discord-Bearer-Token': token,
        },
    })
        .then((res) => res.json())
        .then(
            (json) => json,
            () => ({}),
        );
}
