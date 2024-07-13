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

# Call the get_petition_text function of the smart contract using cast
petition_text=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getPetitionText() (string)")
echo "Petition text: $petition_text"

# Call the get_petition_title function of the smart contract using cast
petition_title=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getPetitionTitle() (string)")
echo "Petition title: $petition_title"

# Call the get_source_links function of the smart contract using cast
source_links=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getSourceLinks() (string[])")
echo "Source links: $source_links"

# Call the get_signers_count function of the smart contract using cast
signers_count=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getSignersCount() (uint32)")
echo "Signers count: $signers_count"
