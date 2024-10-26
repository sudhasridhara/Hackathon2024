package simulation

import (
	"math/rand"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/keeper"
	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddAccountInfo(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddAccountInfo{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddAccountInfo simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "AddAccountInfo simulation not implemented"), nil, nil
	}
}
