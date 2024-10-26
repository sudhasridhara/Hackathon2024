package keeper

import (
	"context"
	"encoding/hex"
	"log"
	"time"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashicorp/vault/shamir"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSecretKey(goCtx context.Context, req *types.QueryGetSecretKeyRequest) (*types.QueryGetSecretKeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Process the query
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
		if account.Owner == req.AccountOwner {
			accountInfo = account
		}
	}

	var shareKeys [][]byte
	for _, recoveryAddress := range accountInfo.MultiSign {
		// Convert hex sharekey string to byte slice
		if recoveryAddress.Sharekey != "empty" {
			shareKeyBytes, err := hex.DecodeString(recoveryAddress.Sharekey)
			if err != nil {
				return nil, errorsmod.Wrap(err, "invalid sharekey format")
			}
			shareKeys = append(shareKeys, shareKeyBytes)
		}
	}
	if len(shareKeys) < int(accountInfo.Minshares) {
		return nil, nil
	}

	reconstructedSecret, err := shamir.Combine(shareKeys)
	if err != nil {
		log.Fatalf("Error reconstructing secret: %v", err)
		return nil, errorsmod.Wrap(err, "failed to reconstruct the secret")
	}

	account := types.Account{
		Owner:            accountInfo.Owner,
		Totalshares:      accountInfo.Totalshares,
		Minshares:        accountInfo.Minshares,
		Creator:          accountInfo.Creator,
		Createddate:      accountInfo.Createddate,
		Validated:        accountInfo.Validated,
		DeathDate:        accountInfo.DeathDate,
		Lastmodifieddate: time.Now().Format("2006-01-02 15:04:05"),
		Id:               accountInfo.Id,
		MultiSign:        accountInfo.MultiSign,
		EmailId:          accountInfo.EmailId,
		Activated:        true,
		SentKey:          true,
	}
	k.UpdateAccountInfo(ctx, account)

	return &types.QueryGetSecretKeyResponse{SecretKey: string(reconstructedSecret)}, nil
}
