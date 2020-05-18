package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

var (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidEthNonce CodeType = 1
	CodeInvalidEthAddress CodeType = 2
	CodeJSONMarshalling CodeType = 3
	CodeInvalidEthSymbol CodeType = 4
	CodeInvalidClaimType      CodeType = 5
	CodeInvalidEthereumChainID CodeType = 6
	CodeInvalidAmount         CodeType = 7
	CodeInvalidSymbol         CodeType = 8
	CodeInvalidBurnSymbol     CodeType = 9
)

// ErrInvalidEthNonce implements sdk.Error.
func ErrInvalidEthNonce(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidEthNonce, "invalid ethereum nonce provided, must be >= 0")
}

// ErrInvalidEthAddress implements sdk.Error.
func ErrInvalidEthAddress(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidEthAddress, "invalid ethereum address provided, must be a valid hex-encoded Ethereum address")
}

// ErrInvalidChainID implements sdk.Error.
func ErrInvalidEthereumChainID(codespace sdk.CodespaceType, chainID string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidEthereumChainID, fmt.Sprintf("invalid ethereum chain id '%s'", chainID))
}

// ErrJSONMarshalling implements sdk.Error.
func ErrJSONMarshalling(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeJSONMarshalling, "error marshalling JSON for this claim")
}

// ErrInvalidEthSymbol implements sdk.Error.
func ErrInvalidEthSymbol(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidEthSymbol, "invalid symbol provided, symbol \"eth\" must have null address set as token contract address")
}

// ErrInvalidClaimType implements sdk.Error.
func ErrInvalidClaimType() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidClaimType, "invalid claim type provided")
}

// ErrInvalidAmount implements sdk.Error.
func ErrInvalidAmount() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidAmount, "amount must be a valid integer > 0")
}

// ErrInvalidSymbol implements sdk.Error.
func ErrInvalidSymbol() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidSymbol, "symbol must be 1 character or more")
}

// ErrInvalidBurnSymbol implements sdk.Error.
func ErrInvalidBurnSymbol() sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidBurnSymbol, fmt.Sprintf("symbol of token to burn must be in the form %v{ethereumSymbol}", PeggedCoinPrefix))
}
