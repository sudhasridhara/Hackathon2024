package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateMultisite{}

func NewMsgUpdateMultisite(creator string, account string, id uint64, multisignAddress string, multisignShareKey string) *MsgUpdateMultisite {
	return &MsgUpdateMultisite{
		Creator:           creator,
		Account:           account,
		Id:                id,
		MultisignAddress:  multisignAddress,
		MultisignShareKey: multisignShareKey,
	}
}

func (msg *MsgUpdateMultisite) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
