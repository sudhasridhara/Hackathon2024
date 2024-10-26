package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateEmailId{}

func NewMsgUpdateEmailId(creator string, account string, emailid string) *MsgUpdateEmailId {
	return &MsgUpdateEmailId{
		Creator: creator,
		Account: account,
		Emailid: emailid,
	}
}

func (msg *MsgUpdateEmailId) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
