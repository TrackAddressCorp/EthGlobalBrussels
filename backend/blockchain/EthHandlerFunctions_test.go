package blockchain

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var params DeployParams
var result *DeployResult

func TestDeployContract(t *testing.T) {
    ctx := context.Background()

    params = DeployParams{
        PetitionTitle: "Test Petition",
        PetitionText: "This is a test petition",
    }

    var err error
    result, err = testHandler.DeployPetition(
        ctx,
        params,
    )

	fmt.Println("address: ", result.Contract)
    require.NoError(t, err)
    require.NotEmpty(t, result)
    require.NotZero(t, result.Address)
    require.NotZero(t, result.Contract)
    require.NotEmpty(t, result.Transaction)

    fmt.Println("Deployment Information-----------------")
    fmt.Println("Petition Contract Address: ", result.Address.Hex())
    fmt.Println("Petition Contract Transaction: ", result.Transaction)
    fmt.Println("----------------------------------------")
}

func TestSignPetition(t *testing.T) {
    ctx := context.Background()

    params := SignParams{
        PetitionAddress: result.Address,
        WorldID: "TestWorldID",
    }

    err := testHandler.SignPetition(
        ctx,
        params,
    )

    require.NoError(t, err)
    fmt.Println("Petition succesfully signed")
    fmt.Println("----------------------------------------")
}

func TestGetSignerCount(t *testing.T) {
    ctx := context.Background()

    count, err := testHandler.GetSignerCount(
        ctx,
        result.Address,
    )

    require.NoError(t, err)
    require.Equal(t, uint64(1), count)

    fmt.Println("Signer Count 1: ", count)

    TestSignPetition(t)

    count, err = testHandler.GetSignerCount(
        ctx,
        result.Address,
    )

    require.NoError(t, err)
    require.Equal(t, uint64(2), count)

    fmt.Println("Signer Count 2: ", count)
}

func TestGetTitle(t *testing.T) {
    ctx := context.Background()

    title, err := testHandler.GetTitle(
        ctx,
        result.Address,
    )

    require.NoError(t, err)
    require.Equal(t, params.PetitionTitle, title)
}

func TestGetText(t *testing.T) {
    ctx := context.Background()

    text, err := testHandler.GetText(
        ctx,
        result.Address,
    )

    require.NoError(t, err)
    require.Equal(t, params.PetitionText, text)
}
