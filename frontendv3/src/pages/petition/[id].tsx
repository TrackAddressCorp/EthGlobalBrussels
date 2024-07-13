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

    const handleProof = async (result: ISuccessResult) => {
        console.log(result);
        const proofData = {
            merkle_root: result.merkle_root,
            nullifier_hash: result.nullifier_hash,
            proof: result.proof,
            verification_level: result.verification_level,
            action: "1",
          };

          try {
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
          } catch (error) {
            console.error('Error:', error);
            setError(error.message);
          }
    }
    return (
        <ChakraProvider>
            <div>
                <Card padding="10px">
                    <CardHeader>
                        <HStack justify="space-between">
                            <Heading size='md'>{item.title}</Heading>
                            <Text>Signs: {item.signs}</Text>
                            <IDKitWidget
                                action={item.ID.toString()}
                                app_id="app_staging_f324149022608832a0b719539d2a4311"
                                onSuccess={onSuccess}
                                handleVerify={handleProof}
                                verification_level={VerificationLevel.Orb}
                            >
                                {({ open }) => (
                                    <Button colorScheme="green"onClick={open}>
                                        Sign with World ID
                                    </Button>
                                )}
                            </IDKitWidget>
                        </HStack>
                    </CardHeader>
                    <Divider></Divider>
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
