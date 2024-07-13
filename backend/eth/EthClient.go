package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

/**
wrapper around ethclient.Client
*/
type Client struct {
    *ethclient.Client
}

/**
    * @param ctx context.Context
    * @param rpc string
    * @return *Client
    * @return error
*/
func NewClient(ctx context.Context, rpc string) (*Client, error) {
    ethClient, err := connectToEth(
        ctx,
        rpc, // reminder: rpc = os.Getenv("ETH_RPC_URL")
    )
    if err != nil {
        return nil, err
    }
    defer ethClient.Close()

    return &Client{ethClient}, nil
}

/**
    * @param ctx context.Context
    * @param rpc string
    * @return *ethclient.Client
    * @return error
*/
func connectToEth(ctx context.Context, rpc string) (*ethclient.Client, error) {
    client, err := ethclient.DialContext(
        ctx,
        rpc,
    )
    if err != nil {
        return nil, err
    }

    return client, nil
}
