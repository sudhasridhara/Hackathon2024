package keeper

import (
	"context"
	"fmt"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateMultisite(goCtx context.Context, msg *types.MsgUpdateMultisite) (*types.MsgUpdateMultisiteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetAccountInfo(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	updated := false
	for i, pair := range val.MultiSign {
		if pair.Address == msg.MultisignAddress {
			val.MultiSign[i].Sharekey = msg.MultisignShareKey
			updated = true
			break
		}
	}
	if updated {
		k.UpdateAccountInfo(ctx, val)
	}
	return &types.MsgUpdateMultisiteResponse{}, nil
}
