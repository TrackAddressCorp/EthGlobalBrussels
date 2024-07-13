'use client'

import { ArrowBackIcon } from '@chakra-ui/icons';
import { Box, Card, CardBody, CardHeader, Center, ChakraProvider, Heading, HStack, IconButton, Text } from '@chakra-ui/react';
import { useRouter } from 'next/router';

import { useEffect, useState } from 'react';

interface Petition {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    title: string;
    description: string | null;
    finished: boolean;
    signs: number;
    individual_votes: any | null;
    pdfs: Pdf[];
}

interface Pdf {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    petition_id: number;
    pdf_url: string;
}

const HomeButton = () => {
    const router = useRouter();

    return (
        <Box position="fixed" bottom="4" left="4">
            <IconButton
                colorScheme="teal"
                aria-label="Add"
                icon={<ArrowBackIcon />}
                size="lg"
                onClick={() => router.push('/')} // Add onClick event to redirect
            />
        </Box>
    );
};

export default function ItemDetail() {
    const router = useRouter();
    const { id } = router.query;
    const [item, setItem] = useState<Petition | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (!id) {
            return;
        }

        fetch(`http://localhost:4242/petition/${id}`)
            .then(response => {
                if (!response.ok) {
                    if (response.status !== 200) {
                        throw new Error(response.json().status_msg);
                    }
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                setLoading(false);
                setItem(data.petition);
            })
            .catch(error => {
                setError(error.message);
                setLoading(false);
            });
    }, [id]);

    if (loading) {
        return (
            <ChakraProvider>
                <p>Loading...</p>
            </ChakraProvider>
        );
    }

    if (error) {
        return (
            <ChakraProvider>
                <Center>
                    <Text>Error: {error}</Text>
                </Center>
            </ChakraProvider>
        )
    }

    if (!item) {
        return (
            <ChakraProvider>
                <Center>
                    <Text>Item not found.</Text>
                </Center>
            </ChakraProvider>
        );
    }

    return (
        <ChakraProvider>
            <div>
                <Card padding="10px">
                    <CardHeader>
                        <HStack justify="space-between">
                            <Heading size='md'>{item.title}</Heading>
                            <Text>Signs: {item.signs}</Text>
                        </HStack>
                    </CardHeader>
                    <CardBody>
                        <Text pt='2' fontSize='sm' marginBottom="20px">
                            {item.description}
                        </Text>
                        {item.pdfs.map(pdf => (
                            <div key={pdf.ID} style={{ marginBottom: '20px' }}>
                                <iframe
                                    src={pdf.pdf_url}
                                    style={{ width: '100%', height: '500px' }}
                                    allowFullScreen
                                ></iframe>
                            </div>
                        ))}
                    </CardBody>
                </Card>
                <HomeButton />
            </div>
        </ChakraProvider>
    );
};
