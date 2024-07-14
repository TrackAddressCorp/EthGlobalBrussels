#!/bin/bash

# Load variables from .env file
set -o allexport
source scripts/.env.tmp
set +o allexport

# Check for variables
if [ -z "$CONTRACT_ADDRESS" ] 
then
    echo "CONTRACT_ADDRESS is not set"
    echo "You can run the script by setting the variables at the beginning: CONTRACT_ADDRESS=0x sign_petition.sh world_id"
    exit 0
fi

prev_signers_count=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getSignersCount() (uint32)")
cast send --rpc-url $RPC_URL --private-key $PRIVATE_KEY $CONTRACT_ADDRESS "sign(uint256)" "$1"
next_signers_count=$(cast call --rpc-url $RPC_URL $CONTRACT_ADDRESS "getSignersCount() (uint32)")

echo "Signers count BEFORE: $prev_signers_count"
echo "Signed the petition with world_id: $1"
echo "Signers count AFTER : $next_signers_count"
