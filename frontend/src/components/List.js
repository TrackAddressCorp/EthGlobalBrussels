import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

function List() {
    const [items, setItems] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        fetch('https://api.example.com/items')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                setItems(data);
                setLoading(false);
            })
            .catch(error => {
                setError(error);
                setLoading(false);
            });
    }, []);

    if (loading) {
        return <p>Loading...</p>;
    }

    if (error) {
        return <p>Error: {error.message}</p>;
    }

    return (
        <ul>
            {items.map(item => (
                <li key={item.id}>
                    <Link to={`/item/${item.id}`}>{item.name}</Link>
                </li>
            ))}
        </ul>
    );
}

export default List;
