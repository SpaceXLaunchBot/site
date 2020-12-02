export default function GetNextLaunch() {
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
