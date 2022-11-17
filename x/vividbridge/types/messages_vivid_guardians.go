package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateVividGuardians = "create_vivid_guardians"
	TypeMsgUpdateVividGuardians = "update_vivid_guardians"
	TypeMsgDeleteVividGuardians = "delete_vivid_guardians"
)

var _ sdk.Msg = &MsgCreateVividGuardians{}

func NewMsgCreateVividGuardians(creator string, keys []string, expirationTime uint64) *MsgCreateVividGuardians {
	return &MsgCreateVividGuardians{
		Creator:        creator,
		Keys:           keys,
		ExpirationTime: expirationTime,
	}
}

func (msg *MsgCreateVividGuardians) Route() string {
	return RouterKey
}

func (msg *MsgCreateVividGuardians) Type() string {
	return TypeMsgCreateVividGuardians
}

func (msg *MsgCreateVividGuardians) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVividGuardians) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVividGuardians) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateVividGuardians{}

func NewMsgUpdateVividGuardians(creator string, id uint64, keys []string, expirationTime uint64) *MsgUpdateVividGuardians {
	return &MsgUpdateVividGuardians{
		Id:             id,
		Creator:        creator,
		Keys:           keys,
		ExpirationTime: expirationTime,
	}
}

func (msg *MsgUpdateVividGuardians) Route() string {
	return RouterKey
}

func (msg *MsgUpdateVividGuardians) Type() string {
	return TypeMsgUpdateVividGuardians
}

func (msg *MsgUpdateVividGuardians) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateVividGuardians) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateVividGuardians) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteVividGuardians{}

func NewMsgDeleteVividGuardians(creator string, id uint64) *MsgDeleteVividGuardians {
	return &MsgDeleteVividGuardians{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteVividGuardians) Route() string {
	return RouterKey
}

func (msg *MsgDeleteVividGuardians) Type() string {
	return TypeMsgDeleteVividGuardians
}

func (msg *MsgDeleteVividGuardians) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVividGuardians) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVividGuardians) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
