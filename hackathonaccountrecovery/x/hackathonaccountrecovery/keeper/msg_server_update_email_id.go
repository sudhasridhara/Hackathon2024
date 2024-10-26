package keeper

import (
	"context"
	"time"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateEmailId(goCtx context.Context, msg *types.MsgUpdateEmailId) (*types.MsgUpdateEmailIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))

	// Create an iterator for the entire store
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	var val types.Account
	// Iterate through all stored events
	for ; iterator.Valid(); iterator.Next() {
		var account types.Account

		// Unmarshal the event data into the event object
		k.cdc.MustUnmarshal(iterator.Value(), &account)

		// Check if the event's name matches the one we're looking for
		if account.Owner == msg.Account {
			val = account
		}
	}

	account := types.Account{
		Owner:            val.Owner,
		Totalshares:      val.Totalshares,
		Minshares:        val.Minshares,
		Creator:          val.Creator,
		Createddate:      val.Createddate,
		Validated:        val.Validated,
		DeathDate:        val.DeathDate,
		Lastmodifieddate: time.Now().Format("2006-01-02 15:04:05"),
		Id:               val.Id,
		MultiSign:        val.MultiSign,
		EmailId:          msg.Emailid,
		Activated:        true,
		SentKey:          false,
	}
	k.UpdateAccountInfo(ctx, account)

	return &types.MsgUpdateEmailIdResponse{}, nil
}
