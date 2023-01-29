package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendText = "send_text"

var _ sdk.Msg = &MsgSendText{}

func NewMsgSendText(creator string, receiver string, body string) *MsgSendText {
	return &MsgSendText{
		Creator:  creator,
		Receiver: receiver,
		Body:     body,
	}
}

func (msg *MsgSendText) Route() string {
	return RouterKey
}

func (msg *MsgSendText) Type() string {
	return TypeMsgSendText
}

func (msg *MsgSendText) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendText) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendText) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
