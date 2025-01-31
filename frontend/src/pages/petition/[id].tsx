'use client'

import { ArrowBackIcon } from '@chakra-ui/icons';
import { Box, Button, Card, CardBody, CardHeader, Center, ChakraProvider, Divider, Heading, HStack, IconButton, Spinner, Text, useToast, VStack } from '@chakra-ui/react';
import { IDKitWidget, ISuccessResult, VerificationLevel } from '@worldcoin/idkit';
import { useRouter } from 'next/router';
import { useCallback, useEffect, useState } from 'react';


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
    const toast = useToast();

    const fetchData = useCallback(() => {
        if (!id) {
            return;
        }

        setLoading(true);
        setError(null);

        fetch(`http://localhost:4242/petition/${id}`)
            .then(async response => {
                if (!response.ok) {
                    if (response.status !== 200) {
                        const errorData = await response.json();
                        throw new Error(errorData.status_msg);
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

    useEffect(() => {
        if (!error) {
            fetchData();
        }
    }, [fetchData]);

    if (loading) {
        return (
            <ChakraProvider>
                <Box p={5} display="flex" justifyContent="center" alignItems="center" height="100vh">
                    <Spinner size="xl" />
                </Box>
            </ChakraProvider>
        );
    }

    if (error) {
        return (
            <ChakraProvider>
                <Center padding={100}>
                    <Text>Error: {error}</Text>
                </Center>
            </ChakraProvider>
        );
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

    const onSuccess = (data: ISuccessResult) => {
        console.log(data);
        setItem({ ...item, signs: item.signs + 1 });
    }

    const handleProof = async (result: ISuccessResult) => {
        console.log(result);
        const proofData = {
            merkle_root: result.merkle_root,
            nullifier_hash: result.nullifier_hash,
            proof: result.proof,
            verification_level: result.verification_level,
            action: item.ID.toString(),
        };

        const response = await fetch('http://localhost:4242/petition/sign', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(proofData),
        });

        const data = await response.json();
        if (data.status_code !== 200) {
            throw new Error(data.status_msg);
        }
        console.log('Success:', data);
    }

    return (
        <ChakraProvider>
            <Box
                minH="100vh"
                backgroundImage="url('/background.png')" // Path to your background image
                backgroundPosition="center"
                backgroundRepeat="no-repeat"
                backgroundSize="cover"
            >
                <Box maxW="container.lg" mx="auto" p="4">
                    <Card p="4" boxShadow="md" borderRadius="md">
                        <CardHeader>
                            <HStack justify="space-between" alignItems="center">
                                <Heading size="md">{item.title}</Heading>
                                <Text>Signs: {item.signs}</Text>
                                <IDKitWidget
                                    action={item.ID.toString()}
                                    app_id={process.env.NEXT_PUBLIC_WLD_APP_ID as `app_${string}`}
                                    onSuccess={onSuccess}
                                    handleVerify={handleProof}
                                    verification_level={VerificationLevel.Orb}
                                >
                                    {({ open }) => (
                                        <Button colorScheme="green" onClick={open}>
                                            Sign with World ID
                                        </Button>
                                    )}
                                </IDKitWidget>
                            </HStack>
                        </CardHeader>
                        <Divider my="4" />
                        <CardBody>
                            <Text pt="2" fontSize="sm" mb="4">
                                {item.description}
                            </Text>
                            <VStack spacing="4">
                                {item.pdfs.map(pdf => (
                                    <Box key={pdf.ID} w="full" borderRadius="md" overflow="hidden" boxShadow="sm">
                                        <iframe
                                            src={pdf.pdf_url}
                                            style={{ width: '100%', height: '500px', border: 'none' }}
                                            allowFullScreen
                                        />
                                    </Box>
                                ))}
                            </VStack>
                        </CardBody>
                    </Card>
                    <Box mt="4">
                        <HomeButton />
                    </Box>
                </Box>
            </Box>
        </ChakraProvider>
    );
}
