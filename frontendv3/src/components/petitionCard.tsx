import React from 'react';
import { Box, Heading, Text, Flex, Spacer, LinkBox, LinkOverlay } from '@chakra-ui/react';
import { useRouter } from 'next/navigation';

const PetitionCard = ({ id, title, description, signs }: { id: number, title: string, description: string, signs: number }) => {
  const router = useRouter();
  // Function to handle navigation
  const navigateToPetition = () => {
    router.push(`/petition/${id}`);
  };

  return (
    <LinkBox as="article" onClick={navigateToPetition} style={{ cursor: 'pointer' }}> {/* Make the whole card clickable */}
      <Box
        borderWidth="1px"
        borderRadius="lg"
        overflow="hidden"
        p={4}
        boxShadow="md"
        bg="white"
        height="400px" // Fixed height
        overflowY="auto" // Allows scrolling within the card for overflow content
      >
        <Flex justifyContent="space-between" alignItems="center" mb={4}>
          <Heading as="h3" size="lg">
            <LinkOverlay href={`/petition/${id}`}>{title}</LinkOverlay> {/* Use LinkOverlay for better accessibility */}
          </Heading>
          <Text fontSize="sm" color="gray.500">
            {signs} Signs
          </Text>
        </Flex>
        <Text fontSize="md" mb={4}>
          {description}
        </Text>
      </Box>
    </LinkBox>
  );
};

export default PetitionCard;
