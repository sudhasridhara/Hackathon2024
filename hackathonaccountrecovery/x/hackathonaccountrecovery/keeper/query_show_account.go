package keeper

import (
	"context"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowAccount(goCtx context.Context, req *types.QueryShowAccountRequest) (*types.QueryShowAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))

	// Create an iterator for the entire store
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	var accountInfo types.Account
	// Iterate through all stored events
	for ; iterator.Valid(); iterator.Next() {
		var account types.Account

		// Unmarshal the event data into the event object
		k.cdc.MustUnmarshal(iterator.Value(), &account)

		// Check if the event's name matches the one we're looking for
		if account.Owner == req.Owner {
			accountInfo = account
		}
	}

	return &types.QueryShowAccountResponse{Account: &accountInfo}, nil
}
