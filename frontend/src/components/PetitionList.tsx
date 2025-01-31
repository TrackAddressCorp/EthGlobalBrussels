import React, { useEffect, useState } from 'react';
import { ChakraProvider, Box, Spinner, SimpleGrid, Center } from '@chakra-ui/react';
import PetitionCard from './petitionCard';

const PetitionList = () => {
    const [petitions, setPetitions] = useState<{
        signs: number;
        description: string;
        title: string; ID: number
    }[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchPetitions = async () => {
            try {
                const response = await fetch('http://localhost:4242/petitions');
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                let data = await response.json();
                setPetitions(data.petitions);
                setLoading(false);
            } catch (error) {
                console.error('Error fetching petitions:', error);
                setLoading(false);
                setPetitions([]); // Ensure petitions is set to an empty array on error
            }
        };

        fetchPetitions();
    }, []);

    if (loading) {
        return (
            <ChakraProvider>
                <Box p={5} display="flex" justifyContent="center" alignItems="center" height="100vh">
                    <Spinner size="xl" />
                </Box>
            </ChakraProvider>
        );
    }

    return (
        <ChakraProvider>
            <Box p={5}>
                <SimpleGrid columns={{ sm: 1, md: 2, lg: 3 }} spacing={5}>
                    {petitions
                        .sort((a, b) => b.signs - a.signs)
                        .map((petition) => (
                            <PetitionCard
                                id={petition.ID}
                                title={petition.title}
                                description={petition.description}
                                signs={petition.signs}
                            />
                        ))}
                </SimpleGrid>
            </Box>
        </ChakraProvider>
    );
};

export default PetitionList;
