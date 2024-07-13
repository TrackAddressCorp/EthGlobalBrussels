#!/bin/bash

# Load variables from .env file
set -o allexport
source scripts/.env.tmp
set +o allexport

# Check for variables
if [ -z "$CONTRACT_ADDRESS" ] 
then
    echo "CONTRACT_ADDRESS is not set"
    echo "You can run the script by setting the variables at the beginning: CONTRACT_ADDRESS=0x get_source_links.sh"
    exit 0
fi

# Call the get_source_links function of the smart contract using cast
source_links=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getSourceLinks() (string[])")

echo "Source links: $source_links"
