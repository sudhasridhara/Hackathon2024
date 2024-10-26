package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgValidateDeath{}

func NewMsgValidateDeath(creator string, owner string, isvalidated bool, id uint64) *MsgValidateDeath {
	return &MsgValidateDeath{
		Creator:     creator,
		Owner:       owner,
		Isvalidated: isvalidated,
		Id:          id,
	}
}

func (msg *MsgValidateDeath) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
