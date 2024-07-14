'use client'

import { AddIcon } from '@chakra-ui/icons';
import { ChakraProvider, Box, SimpleGrid, Center, Heading, IconButton, keyframes, Button } from '@chakra-ui/react';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import RandomlyMovingPepeMemes from '~/components/pepe';
import PetitionList from '~/components/PetitionList';

let animationsEnabled = false; // Global boolean to control animations

const ToggleAnimationsButton = ({ toggleAnimations }: { toggleAnimations: () => void }) => {
  return (
    <Box position="fixed" bottom="20px" right="20px">
      <Button onClick={toggleAnimations} aria-label="Toggle Animations">
        <img src="/favicon.ico" alt="Toggle Animations" style={{ width: 24, height: 24 }} />
      </Button>
    </Box>
  );
};

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
  const [animationsEnabled, setAnimationsEnabled] = useState(false);

  const toggleAnimations = () => {
    setAnimationsEnabled(!animationsEnabled);
  };


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
      <RandomlyMovingPepeMemes animationsEnabled={animationsEnabled} />
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
            animation={animationsEnabled ? `${pulse} 3s ease-in-out infinite, ${fadeIn} 2s` : 'none'}
          >
            <Heading as="h1" size="2xl" mb={4} color="teal.500">
              Welcome to Pepetition!
            </Heading>
          </Box>
        </Center>
        <PetitionList />
        <PlusButton />
        <ToggleAnimationsButton toggleAnimations={toggleAnimations} />
      </Box>
    </ChakraProvider>
  );
};

export default App;
