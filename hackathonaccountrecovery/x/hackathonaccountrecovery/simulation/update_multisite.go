package simulation

import (
	"math/rand"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/keeper"
	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateMultisite(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateMultisite{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateMultisite simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "UpdateMultisite simulation not implemented"), nil, nil
	}
}
