package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitiateDeathValidation{}

func NewMsgInitiateDeathValidation(creator string, owner string, deathDate string, id uint64) *MsgInitiateDeathValidation {
	return &MsgInitiateDeathValidation{
		Creator:   creator,
		Owner:     owner,
		DeathDate: deathDate,
		Id:        id,
	}
}

func (msg *MsgInitiateDeathValidation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
