package blockchain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var params DeployParams

func TestDeployContract(t *testing.T) {
	ctx := context.Background()

	params = DeployParams{
		PetitionTitle: "Test Petition",
		PetitionText: "This is a test petition",
	}
	result, err := testHandler.DeployPetition(
		ctx,
		params,
	)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotZero(t, result.Address)
	require.NotZero(t, result.Contract)
	require.NotEmpty(t, result.Transaction)
}