package keeper

import (
	"context"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAccount(goCtx context.Context, req *types.QueryListAccountRequest) (*types.QueryListAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))

	var accounts []types.Account
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var account types.Account
		if err := k.cdc.Unmarshal(value, &account); err != nil {
			return err
		}

		accounts = append(accounts, account)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListAccountResponse{Account: accounts, Pagination: pageRes}, nil
}
