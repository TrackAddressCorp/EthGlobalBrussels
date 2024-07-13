
'use client'

import React, { useEffect, useState } from 'react';
import { ChakraProvider, Box, Spinner, SimpleGrid, Center, Heading } from '@chakra-ui/react';
import PetitionList from '~/components/petitionList';

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
    </ChakraProvider>
  );
};

export default App;
