package blockchain_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	blockchain "github.com/TrackAddressCorp/EthGlobalBrussels/backend/blockchain"
)

var testHandler *blockchain.Handler

func TestMain(m *testing.M) {
	var err error
	testHandler, err = blockchain.NewHandler(
		context.Background(),
		os.Getenv("ARBITRUM_RPC"),
	)
	require.NoError(m, err)
	os.Exit(m.Run())
}