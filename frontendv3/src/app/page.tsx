
'use client'

import { AddIcon } from '@chakra-ui/icons';
import { ChakraProvider, Box, Spinner, SimpleGrid, Center, Heading, IconButton } from '@chakra-ui/react';
import PetitionList from '~/components/petitionList';

const PlusButton = () => {
  return (
    <Box position="fixed" bottom="4" left="4">
      <IconButton
        colorScheme="teal"
        aria-label="Add"
        icon={<AddIcon />}
        size="lg"
      />
    </Box>
  );
};

const App = () => {
  return (
    <ChakraProvider>
      <Center>
        <Box p={5}>
          <Heading as="h1" size="xl" mb={4}>
            Welcome to the Petition Platform
          </Heading>
        </Box>
      </Center>
      <PetitionList />
      <PlusButton />
    </ChakraProvider>
  );
};

export default App;
