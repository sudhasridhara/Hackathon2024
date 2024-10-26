package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"
)

func (k Keeper) AppendAccountDetails(ctx sdk.Context, account types.Account) uint64 {
	count := k.GetAccountCount(ctx)
	account.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))
	appendedValue := k.cdc.MustMarshal(&account)
	store.Set(GetAccountIDBytes(account.Id), appendedValue)
	k.SetAccountCount(ctx, count+1)
	return count
}

func (k Keeper) GetAccountCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MemCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetAccountIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetAccountCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MemCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetAccountInfo(ctx sdk.Context, id uint64) (val types.Account, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))
	b := store.Get(GetAccountIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) UpdateAccountInfo(ctx sdk.Context, account types.Account) {
	logger := ctx.Logger()
	logger.Info("Attempting to update event", "event_id", account.Id)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))
	val, found := k.GetAccountInfo(ctx, account.Id)
	if !found {
		logger.Error("Event not found for update", "event_id", val.Id)
		//return fmt.Errorf("event with ID %d does not exist", event.Id)
	}
	b := k.cdc.MustMarshal(&account)
	store.Set(GetAccountIDBytes(account.Id), b)
	logger.Info("Event updated successfully", "event_id", account.Id)
}

func (k Keeper) RemoveAccount(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))
	store.Delete(GetAccountIDBytes(id))
}
