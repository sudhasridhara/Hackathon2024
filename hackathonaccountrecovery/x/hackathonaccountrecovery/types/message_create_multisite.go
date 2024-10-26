package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMultisite{}

func NewMsgCreateMultisite(creator string, account string, id uint64, multisign string) *MsgCreateMultisite {
	return &MsgCreateMultisite{
		Creator:   creator,
		Account:   account,
		Id:        id,
		Multisign: multisign,
	}
}

func (msg *MsgCreateMultisite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
