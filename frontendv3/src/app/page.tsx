
'use client'

import { Link } from '@chakra-ui/next-js'

import { Card, CardHeader, CardBody, CardFooter, Heading, Stack, StackDivider, Box , Text, Center} from '@chakra-ui/react'
export default function Home() {

  return (
    <div>
      <Center>

        <Card>
          <CardHeader>
            <Heading size='md'>Client Report</Heading>
          </CardHeader>

          <CardBody>
            <Stack divider={<StackDivider />} spacing='4'>
              <Box>
                <Heading size='xs' textTransform='uppercase'>
                  Summary
                </Heading>
                <Text pt='2' fontSize='sm'>
                  View a summary of all your clients over the last month.
                </Text>
              </Box>
              <Box>
                <Heading size='xs' textTransform='uppercase'>
                  Overview
                </Heading>
                <Text pt='2' fontSize='sm'>
                  Check out the overview of your clients.
                </Text>
              </Box>
              <Box>
                <Heading size='xs' textTransform='uppercase'>
                  Analysis
                </Heading>
                <Text pt='2' fontSize='sm'>
                  See a detailed analysis of all your business clients.
                </Text>
              </Box>
            </Stack>
          </CardBody>
        </Card>
      </Center>

      <Link href='/item/1' color='blue.400' _hover={{ color: 'blue.500' }}>
        About
      </Link>

    </div>
  );
}
