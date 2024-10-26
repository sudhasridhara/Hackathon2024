package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"hackathonaccountrecovery/x/hackathonaccountrecovery/types"
	"log"
	"time"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMultisite(goCtx context.Context, msg *types.MsgCreateMultisite) (*types.MsgCreateMultisiteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetAccountInfo(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	var recoveryAddresses []types.RecoveryAddresses
	err := json.Unmarshal([]byte(msg.Multisign), &recoveryAddresses)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
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
		Id:               msg.Id,
		MultiSign:        recoveryAddresses,
		EmailId:          val.EmailId,
		Activated:        true,
		SentKey:          false,
	}
	k.UpdateAccountInfo(ctx, account)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateMultisiteResponse{}, nil
}
