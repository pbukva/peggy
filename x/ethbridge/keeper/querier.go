package keeper

import (
	"fmt"
	"github.com/Workiva/go-datastructures/threadsafe/err"
	"strconv"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/peggy/x/ethbridge/types"
	oracletypes "github.com/cosmos/peggy/x/oracle/types"
)

// TODO: move to x/oracle

// NewQuerier is the module level router for state queries
func NewQuerier(keeper types.OracleKeeper, cdc *codec.Codec, codespace sdk.CodespaceType) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryEthProphecy:
			return queryEthProphecy(ctx, cdc, req, keeper, codespace)
		default:
			return nil, sdk.ErrUnknownRequest("unknown ethbridge query endpoint")
		}
	}
}

func queryEthProphecy(
	ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery, keeper types.OracleKeeper, codespace sdk.CodespaceType) ([]byte, sdk.Error) {
	var params types.QueryEthProphecyParams

	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to parse params: %s", err.Error()))
	}

	id := strconv.Itoa(params.EthereumChainID) + strconv.Itoa(params.Nonce) + params.EthereumSender.String()
	prophecy, found := keeper.GetProphecy(ctx, id)
	if !found {
		return nil, oracletypes.ErrProphecyNotFound(codespace, id)
	}

	bridgeClaims, err := types.MapOracleClaimsToEthBridgeClaims(
		params.EthereumChainID, params.BridgeContractAddress, params.Nonce, params.Symbol, params.TokenContractAddress,
		params.EthereumSender, prophecy.ValidatorClaims, types.CreateEthClaimFromOracleString)
	if err != nil {
		return nil, err
	}

	response := types.NewQueryEthProphecyResponse(prophecy.ID, prophecy.Status, bridgeClaims)

	return cdc.MarshalJSONIndent(response, "", "  ")
}
