#!/bin/bash

# Load variables from .env file
set -o allexport
source scripts/.env.tmp
set +o allexport

# Check for variables
if [ -z "$CONTRACT_ADDRESS" ] 
then
    echo "CONTRACT_ADDRESS is not set"
    echo "You can run the script by setting the variables at the beginning: CONTRACT_ADDRESS=0x init_petition.sh petition_text petition_title"
    exit 0
fi


# Call the init function of the smart contract using cast
cast send --rpc-url $RPC_URL $CONTRACT_ADDRESS --private-key $PRIVATE_KEY "init(string, string, string[])" "$1" "$2" "[$3, $4]"

echo "Initialized petition with:"
echo "  - text: $1"
echo "  - title: $2"
echo " - links: $3, $4"