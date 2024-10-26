package keeper

import (
	"context"
	"time"

	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAccountInfo(goCtx context.Context, msg *types.MsgAddAccountInfo) (*types.MsgAddAccountInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var account = types.Account{
		Owner:            msg.Owner,
		Totalshares:      msg.Totalshares,
		Minshares:        msg.Minshares,
		Creator:          msg.Creator,
		Createddate:      time.Now().Format("2006-01-02 15:04:05"),
		Lastmodifieddate: time.Now().Format("2006-01-02 15:04:05"),
		DeathDate:        "",
		Validated:        false,
		EmailId:          msg.EmailId,
		SentKey:          false,
	}
	id := k.AppendAccountDetails(
		ctx,
		account,
	)
	if id == 0 {
		var account = types.Account{
			Owner:            msg.Owner,
			Totalshares:      msg.Totalshares,
			Minshares:        msg.Minshares,
			Creator:          msg.Creator,
			Createddate:      time.Now().Format("2006-01-02 15:04:05"),
			Lastmodifieddate: time.Now().Format("2006-01-02 15:04:05"),
			DeathDate:        "",
			Validated:        false,
			Id:               id,
		}
		k.UpdateAccountInfo(
			ctx,
			account,
		)

	}

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddAccountInfoResponse{Id: id}, nil
}
