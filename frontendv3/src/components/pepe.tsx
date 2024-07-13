import React, { useEffect, useState } from 'react';

// Assuming you have an array of Pepe meme images
const pepeMemes = [
    "pepe1.jpg", // Update with actual path to pepe memes
    "pepe2.jpg",
    "pepe3.jpg"
];

const RandomlyMovingPepeMemes = () => {
    // Initialize state with random positions, rotations, and scales for each meme
    const [styles, setStyles] = useState(pepeMemes.map(() => ({
        top: `calc(${Math.random() * 100}% - 50px)`,
        left: `calc(${Math.random() * 100}% - 50px)`,
        transform: `rotate(${Math.random() * 360}deg) scale(${0.5 + Math.random() * 1.5})`,
    })));

    useEffect(() => {
        // Function to update styles
        const updateStyles = () => {
            setStyles(styles.map(() => ({
                top: `calc(${Math.random() * 100}% - 50px)`,
                left: `calc(${Math.random() * 100}% - 50px)`,
                transform: `rotate(${Math.random() * 360}deg) scale(${0.5 + Math.random() * 1.5})`,
            })));
        };

        // Set interval to update styles every cycle (e.g., 1000ms)
        const intervalId = setInterval(updateStyles, 1000);

        // Cleanup interval on component unmount
        return () => clearInterval(intervalId);
    }, []); // Empty dependency array means this effect runs once on mount

    return (
        <>
            {pepeMemes.map((memeSrc, index) => (
                <img
                    key={index}
                    src={memeSrc}
                    alt={`Pepe Meme ${index + 1}`}
                    style={{
                        position: 'absolute',
                        ...styles[index],
                        width: '100px', // Set a fixed size or make it dynamic as per your requirement
                        height: '100px',
                        transition: 'transform 1s, top 1s, left 1s', // Smooth transition for transformations
                    }}
                />
            ))}
        </>
    );
};

export default RandomlyMovingPepeMemes;
