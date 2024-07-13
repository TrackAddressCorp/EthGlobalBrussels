package blockchain

import (
	"context"
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

    require.NoError(t, err)
    require.NotEmpty(t, result)
    require.NotZero(t, result.Address)
    require.NotZero(t, result.Contract)
    require.NotEmpty(t, result.Transaction)
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
}

func TestGetSignerCount(t *testing.T) {
    ctx := context.Background()

    count, err := testHandler.GetSignerCount(
        ctx,
        result.Address,
    )

    require.NoError(t, err)
    require.NotZero(t, count)
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
