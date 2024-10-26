package keeper

import (
	"context"
	"fmt"
	"time"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) InitiateDeathValidation(goCtx context.Context, msg *types.MsgInitiateDeathValidation) (*types.MsgInitiateDeathValidationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetAccountInfo(ctx, msg.Id)

	account := types.Account{
		Owner:            val.Owner,
		Totalshares:      val.Totalshares,
		Minshares:        val.Minshares,
		Creator:          val.Creator,
		Createddate:      val.Createddate,
		Validated:        val.Validated,
		DeathDate:        msg.DeathDate,
		Lastmodifieddate: time.Now().Format("2006-01-02 15:04:05"),
		Id:               msg.Id,
		MultiSign:        val.MultiSign,
		EmailId:          val.EmailId,
		Activated:        val.Activated,
		SentKey:          false,
	}

	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.UpdateAccountInfo(ctx, account)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgInitiateDeathValidationResponse{}, nil
}
