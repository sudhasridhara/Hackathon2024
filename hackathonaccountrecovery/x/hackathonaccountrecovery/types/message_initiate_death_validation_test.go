package types

import (
	"testing"

	"hackathonaccountrecovery/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgInitiateDeathValidation_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInitiateDeathValidation
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInitiateDeathValidation{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgInitiateDeathValidation{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
