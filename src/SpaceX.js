export function GetLaunch(launchNum) {
    return fetch('https://api.spacexdata.com/v4/launches/query', {
        method: 'POST',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            query: { flight_number: launchNum },
            options: {
                limit: 1,
            },
        }),
    })
        .then((res) => res.json())
        .then(
            (json) => json.docs[0],
            () => ({}),
        );
}

export function GetNextLaunch() {
    return fetch('https://api.spacexdata.com/v4/launches/next', {
        method: 'GET',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
        },
    })
        .then((res) => res.json())
        .then(
            (json) => json,
            () => ({}),
        );
}
