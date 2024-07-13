'use client'

import { AddIcon } from '@chakra-ui/icons';
import { ChakraProvider, Box, Spinner, SimpleGrid, Center, Heading, IconButton, Image, keyframes } from '@chakra-ui/react';
import { useRouter } from 'next/navigation';
import RandomlyMovingPepeMemes from '~/components/pepe';
import PetitionList from '~/components/petitionList';

const PlusButton = () => {
  const router = useRouter();

  return (
    <Box position="fixed" bottom="4" left="4">
      <IconButton
        colorScheme="teal"
        aria-label="Add"
        icon={<AddIcon />}
        size="lg"
        onClick={() => router.push('/create')} // Add onClick event to redirect
      />
    </Box>
  );
};

const App = () => {
  // Define keyframes for animations
  const pulse = keyframes`
    0% {
      transform: scale(1);
    }
    50% {
      transform: scale(1.05);
    }
    100% {
      transform: scale(1);
    }
  `;

  const fadeIn = keyframes`
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  `;

  return (
    <ChakraProvider>
      <RandomlyMovingPepeMemes />
      <Box
        minH="100vh"
        backgroundImage="url('/background.png')" // Path to your background image
        backgroundPosition="center"
        backgroundRepeat="no-repeat"
        backgroundSize="cover"
      >
        <Center>
          <Box
            p={5}
            bg="whiteAlpha.800"
            borderRadius="md"
            boxShadow="lg"
            mt={8}
            mb={8}
            textAlign="center"
            position="relative"
            animation={`${pulse} 3s ease-in-out infinite, ${fadeIn} 2s`}
          >
            <Heading as="h1" size="2xl" mb={4} color="teal.500">
              Welcome to the Petition Platform
            </Heading>
          </Box>
        </Center>
        <PetitionList />
        <PlusButton />
      </Box>
    </ChakraProvider>
  );
};

export default App;
