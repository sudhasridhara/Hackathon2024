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

func (k Keeper) GetMappedAccounts(goCtx context.Context, req *types.QueryGetMappedAccountsRequest) (*types.QueryGetMappedAccountsResponse, error) {
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
		for _, recoveryAddress := range account.MultiSign {
			if recoveryAddress.Address == req.Account {
				accounts = append(accounts, account)
				break // Stop iterating once a match is found for this account
			}
		}
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetMappedAccountsResponse{Account: accounts, Pagination: pageRes}, nil
}
