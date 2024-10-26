package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddAccountInfo{}

func NewMsgAddAccountInfo(creator string, owner string, totalshares uint64, minshares uint64, emailId string) *MsgAddAccountInfo {
	return &MsgAddAccountInfo{
		Creator:     creator,
		Owner:       owner,
		Totalshares: totalshares,
		Minshares:   minshares,
		EmailId:     emailId,
	}
}

func (msg *MsgAddAccountInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
