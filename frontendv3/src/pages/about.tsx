'use client'

import { Button, Card, CardHeader, Heading } from "@chakra-ui/react";
import { VerificationLevel, IDKitWidget } from "@worldcoin/idkit";
import type { ISuccessResult } from "@worldcoin/idkit";
import { ChakraProvider } from "@chakra-ui/react";

function onSuccess(data: ISuccessResult) {
  console.log(data);
}

const handleProof = async (result: ISuccessResult) => {
  console.log(result);
}

export default function About() {
  return (
    <ChakraProvider>
    <div>
      <Card>
      <CardHeader>
            <Heading size='md'>Client Report</Heading>
          </CardHeader>
      </Card>
      <IDKitWidget
        action={process.env.NEXT_PUBLIC_WLD_ACTION!}
        app_id={process.env.NEXT_PUBLIC_WLD_APP_ID as `app_${string}`}
        onSuccess={onSuccess}
        handleVerify={handleProof}
        verification_level={VerificationLevel.Orb}
      >
        {({ open }) => (
          <Button className="border border-black rounded-md" onClick={open}>
            <div className="mx-3 my-1">Sign with World ID</div>
          </Button>
        )}
      </IDKitWidget>

    </div>
    </ChakraProvider>
  );
}
