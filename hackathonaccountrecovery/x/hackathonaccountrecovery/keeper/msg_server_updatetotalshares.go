package keeper

import (
	"context"
	"fmt"
	"time"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Updatetotalshares(goCtx context.Context, msg *types.MsgUpdatetotalshares) (*types.MsgUpdatetotalsharesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetAccountInfo(ctx, msg.Id)
	account := types.Account{
		Owner:            val.Owner,
		Totalshares:      msg.Totalshares,
		Minshares:        msg.Minshares,
		Creator:          val.Creator,
		Createddate:      val.Createddate,
		Lastmodifieddate: time.Now().Format("2006-01-02 15:04:05"),
		Validated:        val.Validated,
		DeathDate:        val.DeathDate,
		Id:               msg.Id,
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

	return &types.MsgUpdatetotalsharesResponse{Id: account.Id}, nil
}
