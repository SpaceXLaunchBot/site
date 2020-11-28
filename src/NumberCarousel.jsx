import React from 'react';

export default function NumberCarousel(props) {
    const { number, setNumber } = props;
    const prev2 = number - 2;
    const prev = number - 1;
    const next = number + 1;
    const next2 = number + 2;
    return (
        <ol className="numberCarousel">
            <li><button type="button" onClick={() => setNumber(prev)}>&lt;</button></li>
            <li><button type="button" onClick={() => setNumber(prev2)}>{prev2}</button></li>
            <li><button type="button" onClick={() => setNumber(prev)}>{prev}</button></li>
            <li><button type="button" id="centralButton">{number}</button></li>
            <li><button type="button" onClick={() => setNumber(next)}>{next}</button></li>
            <li><button type="button" onClick={() => setNumber(next2)}>{next2}</button></li>
            <li><button type="button" onClick={() => setNumber(next)}>&gt;</button></li>
        </ol>
    );
}
