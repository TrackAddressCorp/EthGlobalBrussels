#!/bin/bash

# Load variables from .env file
set -o allexport
source scripts/.env.tmp
set +o allexport

# Check for variables
if [ -z "$CONTRACT_ADDRESS" ] 
then
    echo "CONTRACT_ADDRESS is not set"
    echo "You can run the script by setting the variables at the beginning: CONTRACT_ADDRESS=0x get_petition_text.sh"
    exit 0
fi

# Call the get_petition_text function of the smart contract using cast
petition_text=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getPetitionText() (string)")

echo "Petition text: $petition_text"
