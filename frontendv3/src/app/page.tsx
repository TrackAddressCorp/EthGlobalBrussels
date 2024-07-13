'use client'

import { AddIcon } from '@chakra-ui/icons';
import { ChakraProvider, Box, Spinner, SimpleGrid, Center, Heading, IconButton } from '@chakra-ui/react';
import { useRouter } from 'next/navigation';
import PetitionList from '~/components/petitionList';
 // Import useHistory from react-router-dom

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
  return (
    <ChakraProvider>
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
