#!/bin/bash

# Load variables from .env file
set -o allexport
source scripts/.env.tmp
set +o allexport

# Check for variables
if [ -z "$CONTRACT_ADDRESS" ] 
then
    echo "CONTRACT_ADDRESS is not set"
    echo "You can run the script by setting the variables at the beginning: CONTRACT_ADDRESS=0x get_petition_title.sh"
    exit 0
fi

# Call the get_petition_title function of the smart contract using cast
petition_title=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getPetitionTitle() (string)")

echo "Petition title: $petition_title"
