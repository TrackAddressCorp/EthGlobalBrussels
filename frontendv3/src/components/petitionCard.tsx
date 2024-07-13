import React from 'react';
import { Box, Heading, Text } from '@chakra-ui/react';

const PetitionCard = ({ key, title, description }: {key :number,  title: string, description: string }) => {
  return (
    <Box
      borderWidth="1px"
      borderRadius="lg"
      overflow="hidden"
      p={4}
      boxShadow="md"
      bg="white"
      key={key}
      width={{ base: '100%', md: 'auto' }}
    >
      <Heading as="h3" size="lg" mb={4}>
        {title}
      </Heading>
      <Text fontSize="md">
        {description}
      </Text>
    </Box>
  );
};

export default PetitionCard;
