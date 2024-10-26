package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdatetotalshares{}

func NewMsgUpdatetotalshares(creator string, owner string, totalshares uint64, minshares uint64, id uint64) *MsgUpdatetotalshares {
	return &MsgUpdatetotalshares{
		Creator:     creator,
		Owner:       owner,
		Totalshares: totalshares,
		Minshares:   minshares,
		Id:          id,
	}
}

func (msg *MsgUpdatetotalshares) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
