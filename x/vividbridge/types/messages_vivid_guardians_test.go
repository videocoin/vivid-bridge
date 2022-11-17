package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"vivid-bridge/testutil/sample"
)

func TestMsgCreateVividGuardians_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateVividGuardians
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVividGuardians{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateVividGuardians{
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

func TestMsgUpdateVividGuardians_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateVividGuardians
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateVividGuardians{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateVividGuardians{
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

func TestMsgDeleteVividGuardians_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteVividGuardians
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteVividGuardians{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteVividGuardians{
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
