package blockchain

import (
    "context"
    "testing"

    "github.com/stretchr/testify/require"
)

func exampleHandler(t *testing.T) (*Handler){
    // rpc := os.Getenv("ARBITRUM_RPC")
    handler, err := NewHandler(
        context.Background(),
        "https://sepolia-rollup.arbitrum.io/rpc",
        "c7fb5814f19cd9d6f9a69e79a25b35fb29d21a6d02948f0756def9bba6e14af5",
    )

    require.NoError(t, err)
    require.NotEmpty(t, handler)

    require.NotZero(t, handler.Client)
    require.NotZero(t, handler.PrivateKey)
    require.NotZero(t, handler.PublicKey)
    // require.Equal(t, os.Getenv("PUBLIC_KEY"), handler.PublicKeyString)
    require.NotZero(t, handler.PublicAddress)

    return handler
}

func TestNewHandler(t *testing.T) {
    exampleHandler(t)
}
