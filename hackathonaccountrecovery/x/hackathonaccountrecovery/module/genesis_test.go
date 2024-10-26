package hackathonaccountrecovery_test

import (
	"testing"

	keepertest "hackathonaccountrecovery/testutil/keeper"
	"hackathonaccountrecovery/testutil/nullify"
	hackathonaccountrecovery "hackathonaccountrecovery/x/hackathonaccountrecovery/module"
	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HackathonaccountrecoveryKeeper(t)
	hackathonaccountrecovery.InitGenesis(ctx, k, genesisState)
	got := hackathonaccountrecovery.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
