import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { IDKitWidget } from '@worldcoin/idkit';

function ItemDetail() {
    const { id } = useParams();
    const [item, setItem] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        fetch(`http://localhost:4242/petition/${id}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                setItem(data);
                setLoading(false);
            })
            .catch(error => {
                setError(error);
                setLoading(false);
            });
    }, [id]);

    const onSuccess = (result) => {
        console.log('Verification successful:', result);
        // Handle successful verification here
    };

    const onError = (error) => {
        console.error('Verification error:', error);
        // Handle error here
    };

    if (loading) {
        return <p>Loading...</p>;
    }

    if (error) {
        return <p>Error: {error.message}</p>;
    }

    return (
        <div>
            <h1>{item.name}</h1>
            <p>{item.description}</p>
            <IDKitWidget
                app_id="app_GBkZ1KlVUdFTjeMXKlVUdFT" // Replace with your actual app ID
                action="view_item" // Define your action name
                onSuccess={onSuccess}
                onError={onError}
                verification_level="device" // Define your desired verification level
            >
                {({ open }) => <button onClick={open}>Verify with World ID</button>}
            </IDKitWidget>
        </div>
    );
}

export default ItemDetail;
