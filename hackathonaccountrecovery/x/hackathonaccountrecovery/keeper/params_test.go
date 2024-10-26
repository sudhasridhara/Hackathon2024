package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "hackathonaccountrecovery/testutil/keeper"
	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.HackathonaccountrecoveryKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
