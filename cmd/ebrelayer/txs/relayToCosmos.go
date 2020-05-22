package txs

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/peggy/x/ethbridge"
	"github.com/cosmos/peggy/x/ethbridge/types"
)

// RelayToCosmos applies validator's signature to an EthBridgeClaim message containing
// information about an event on the Ethereum blockchain before relaying to the Bridge
func RelayToCosmos(cdc *codec.Codec, moniker string, claim *types.EthBridgeClaim, cliCtx context.CLIContext,
	txBldr authtypes.TxBuilder) error {
	// Packages the claim as a Tendermint message
	msg := ethbridge.NewMsgCreateEthBridgeClaim(*claim)

	err := msg.ValidateBasic()
	if err != nil {
		return err
	}

	// Prepare tx
	txBldr, err2 := utils.PrepareTxBuilder(txBldr, cliCtx)
	if err2 != nil {
		return err2
	}

	// Build and sign the transaction
	txBytes, err2 := txBldr.BuildAndSign(moniker, keys.DefaultKeyPass, []sdk.Msg{msg})
	if err2 != nil {
		return err2
	}

	// Broadcast to a Tendermint node
	res, err2 := cliCtx.BroadcastTxSync(txBytes)
	if err2 != nil {
		return err2
	}

	if err2 = cliCtx.PrintOutput(res); err2 != nil {
		return err2
	}
	return nil
}
