package blockchain

import (
	"context"
	"log"
	"os"
	"testing"
)

var testHandler *Handler

func TestMain(m *testing.M) {
    var err error
    testHandler, err = NewHandler(
        context.Background(),
        // "https://sepolia-rollup.arbitrum.io/rpc",
        // "c7fb5814f19cd9d6f9a69e79a25b35fb29d21a6d02948f0756def9bba6e14af5",
		"HTTP://host.docker.internal:7545",
		"416aaf560788b841a577de5493c95056e789d1945f61912bd5a8fb1e237ad780",
    )
    if err != nil {
        log.Fatal(err)
    }
    os.Exit(m.Run())
}
// 900000000000000000
// 125000000000000000