import React from 'react';

export default function NumberCarousel(props) {
    const { number, setNumber } = props;

    const setter = (e) => {
        const num = parseInt(e.target.textContent, 10);
        if (num > 0) {
            setNumber(num);
        }
    };

    return (
        <ol>
            <li>&lt;</li>
            <li><button type="button" onClick={setter}>{number - 1}</button></li>
            <li><button type="button">{number}</button></li>
            <li><button type="button" onClick={setter}>{number + 1}</button></li>
            <li>&gt;</li>
        </ol>
    );
}
