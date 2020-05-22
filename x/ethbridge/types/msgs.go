package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	gethCommon "github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgLock defines a message for locking coins and triggering a related event
type MsgLock struct {
	CosmosSender     sdk.AccAddress  `json:"cosmos_sender" yaml:"cosmos_sender"`
	Amount           int64           `json:"amount" yaml:"amount"`
	Symbol           string          `json:"symbol" yaml:"symbol"`
	EthereumChainID  int             `json:"ethereum_chain_id" yaml:"ethereum_chain_id"`
	EthereumReceiver EthereumAddress `json:"ethereum_receiver" yaml:"ethereum_receiver"`
}

// NewMsgLock is a constructor function for MsgLock
func NewMsgLock(
	ethereumChainID int, cosmosSender sdk.AccAddress,
	ethereumReceiver EthereumAddress, amount int64, symbol string) MsgLock {
	return MsgLock{
		EthereumChainID:  ethereumChainID,
		CosmosSender:     cosmosSender,
		EthereumReceiver: ethereumReceiver,
		Amount:           amount,
		Symbol:           symbol,
	}
}

// Route should return the name of the module
func (msg MsgLock) Route() string { return RouterKey }

// Type should return the action
func (msg MsgLock) Type() string { return "lock" }

// ValidateBasic runs stateless checks on the message
func (msg MsgLock) ValidateBasic() sdk.Error {
	//TODO(pb): Does this check make sense at all?
	if strconv.Itoa(msg.EthereumChainID) == "" {
		return ErrInvalidEthereumChainID(DefaultCodespace, strconv.Itoa(msg.EthereumChainID))
	}

	if msg.CosmosSender.Empty() {
		return sdk.ErrInvalidAddress("Empty cosmos sender address string")
	}

	if msg.EthereumReceiver.String() == "" {
		return ErrInvalidEthAddress("Empty ethereum receiver address string")
	}

	if !gethCommon.IsHexAddress(msg.EthereumReceiver.String()) {
		return ErrInvalidEthAddress(fmt.Sprintf("Ethereum receiver: %s", msg.EthereumReceiver.String()))
	}

	if msg.Amount <= 0 {
		return ErrInvalidAmount()
	}

	if len(msg.Symbol) == 0 {
		return ErrInvalidSymbol()
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgLock) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgLock) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.CosmosSender}
}

// MsgBurn defines a message for burning coins and triggering a related event
type MsgBurn struct {
	CosmosSender     sdk.AccAddress  `json:"cosmos_sender" yaml:"cosmos_sender"`
	Amount           int64           `json:"amount" yaml:"amount"`
	Symbol           string          `json:"symbol" yaml:"symbol"`
	EthereumChainID  int             `json:"ethereum_chain_id" yaml:"ethereum_chain_id"`
	EthereumReceiver EthereumAddress `json:"ethereum_receiver" yaml:"ethereum_receiver"`
}

// NewMsgBurn is a constructor function for MsgBurn
func NewMsgBurn(
	ethereumChainID int, cosmosSender sdk.AccAddress,
	ethereumReceiver EthereumAddress, amount int64, symbol string) MsgBurn {
	return MsgBurn{
		EthereumChainID:  ethereumChainID,
		CosmosSender:     cosmosSender,
		EthereumReceiver: ethereumReceiver,
		Amount:           amount,
		Symbol:           symbol,
	}
}

// Route should return the name of the module
func (msg MsgBurn) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBurn) Type() string { return "burn" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBurn) ValidateBasic() sdk.Error {
	//TODO(pb): Does this check make sense at all?
	if strconv.Itoa(msg.EthereumChainID) == "" {
		return ErrInvalidEthereumChainID(DefaultCodespace, strconv.Itoa(msg.EthereumChainID))
	}

	if msg.CosmosSender.Empty() {
		return sdk.ErrInvalidAddress("Empty cosmos sender address string")
	}

	if msg.EthereumReceiver.String() == "" {
		return ErrInvalidEthAddress("Empty ethereum receiver address string")
	}

	if !gethCommon.IsHexAddress(msg.EthereumReceiver.String()) {
		return ErrInvalidEthAddress(fmt.Sprintf("Ethereum receiver: %s", msg.EthereumReceiver.String()))	}

	if msg.Amount <= 0 {
		return ErrInvalidAmount()
	}
	prefixLength := len(PeggedCoinPrefix)
	if len(msg.Symbol) <= prefixLength+1 {
		return ErrInvalidBurnSymbol()
	}
	symbolPrefix := msg.Symbol[:prefixLength]
	if symbolPrefix != PeggedCoinPrefix {
		return ErrInvalidBurnSymbol()
	}
	symbolSuffix := msg.Symbol[prefixLength:]
	if len(symbolSuffix) == 0 {
		return ErrInvalidBurnSymbol()
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBurn) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.CosmosSender}
}

// MsgCreateEthBridgeClaim defines a message for creating claims on the ethereum bridge
type MsgCreateEthBridgeClaim EthBridgeClaim

// NewMsgCreateEthBridgeClaim is a constructor function for MsgCreateBridgeClaim
func NewMsgCreateEthBridgeClaim(ethBridgeClaim EthBridgeClaim) MsgCreateEthBridgeClaim {
	return MsgCreateEthBridgeClaim(ethBridgeClaim)
}

// Route should return the name of the module
func (msg MsgCreateEthBridgeClaim) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateEthBridgeClaim) Type() string { return "create_bridge_claim" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateEthBridgeClaim) ValidateBasic() sdk.Error {
	if msg.CosmosReceiver.Empty() {
		return sdk.ErrInvalidAddress("Empty cosmos receiver address string")
	}

	if msg.ValidatorAddress.Empty() {
		return sdk.ErrInvalidAddress("Empty cosmos validator address string")
	}

	if msg.Nonce < 0 {
		return ErrInvalidEthNonce(DefaultCodespace)
	}

	if !gethCommon.IsHexAddress(msg.EthereumSender.String()) {
		return ErrInvalidEthAddress(fmt.Sprintf("Ethereum sender address: %s", msg.EthereumSender.String()))
	}
	if !gethCommon.IsHexAddress(msg.BridgeContractAddress.String()) {
		return ErrInvalidEthAddress(fmt.Sprintf("Ethereum Bridge Contract address: %s", msg.BridgeContractAddress.String()))
	}
	if !gethCommon.IsHexAddress(msg.TokenContractAddress.String()) {
		return ErrInvalidEthAddress(fmt.Sprintf("Ethereum Token Contract address: %s", msg.TokenContractAddress.String()))
	}
	if strings.ToLower(msg.Symbol) == "eth" &&
		msg.TokenContractAddress != NewEthereumAddress("0x0000000000000000000000000000000000000000") {
		return ErrInvalidEthSymbol(DefaultCodespace)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateEthBridgeClaim) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgCreateEthBridgeClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddress)}
}

// MapOracleClaimsToEthBridgeClaims maps a set of generic oracle claim data into EthBridgeClaim objects
func MapOracleClaimsToEthBridgeClaims(
	ethereumChainID int, bridgeContract EthereumAddress, nonce int, symbol string,
	tokenContract EthereumAddress, ethereumSender EthereumAddress,
	oracleValidatorClaims map[string]string,
	f func(int, EthereumAddress, int, EthereumAddress, sdk.ValAddress, string,
	) (EthBridgeClaim, sdk.Error),
) ([]EthBridgeClaim, sdk.Error) {
	mappedClaims := make([]EthBridgeClaim, len(oracleValidatorClaims))
	i := 0
	for validatorBech32, validatorClaim := range oracleValidatorClaims {
		validatorAddress, parseErr := sdk.ValAddressFromBech32(validatorBech32)
		if parseErr != nil {
			return nil, sdk.ErrInvalidAddress(fmt.Sprintf("failed to parse claim: %s", parseErr))
		}
		mappedClaim, err := f(
			ethereumChainID, bridgeContract, nonce, ethereumSender, validatorAddress, validatorClaim)
		if err != nil {
			return nil, err
		}
		mappedClaims[i] = mappedClaim
		i++
	}
	return mappedClaims, nil
}
