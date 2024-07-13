package blockchain_test

import (
	"context"
	"log"
	"os"
	"testing"

	blockchain "github.com/TrackAddressCorp/EthGlobalBrussels/backend/blockchain"
)

var testHandler *blockchain.Handler

func TestMain(m *testing.M) {
	var err error
	testHandler, err = blockchain.NewHandler(
		context.Background(),
		os.Getenv("ARBITRUM_RPC"),
	)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}