#!/bin/bash

# Load variables from .env file
set -o allexport
source scripts/.env.tmp
set +o allexport

# Check for variables
if [ -z "$CONTRACT_ADDRESS" ] 
then
    echo "CONTRACT_ADDRESS is not set"
    echo "You can run the script by setting the variables at the beginning: CONTRACT_ADDRESS=0x get_signers_count.sh"
    exit 0
fi

# Call the get_signers_count function of the smart contract using cast
signers_count=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getSignersCount() (uint32)")

echo "Signers count: $signers_count"
