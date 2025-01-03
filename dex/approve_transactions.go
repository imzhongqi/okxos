package dex

import (
	"context"

	"github.com/imzhongqi/okxos/errcode"
)

type ApproveTransactionsRequest struct {
	// ChainId is the chain ID (e.g., 1 for Ethereum. See Chain IDs)
	ChainId string `json:"chainId"`
	// TokenAddress is the contract address of a token to be sold (e.g., 0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee)
	TokenAddress string `json:"tokenAddress"`
	// ApproveAmount is the amount of token that needs to be permitted (set in minimal divisible units,
	// e.g., 1.00 USDT set as 1000000, 1.00 DAI set as 1000000000000000000)
	ApproveAmount string `json:"approveAmount"`
}

type ApproveTransactionsResult struct {
	// Data is the call data
	Data string `json:"data"`
	// DexContractAddress is the contract address of OKX DEX approve (e.g., 0x6f9ffea7370310cd0f890dfde5e0e061059dcfd9)
	DexContractAddress string `json:"dexContractAddress"`
	// GasLimit is the gas limit (e.g., 50000)
	GasLimit string `json:"gasLimit"`
	// GasPrice is the gas price in wei (e.g., 110000000)
	GasPrice string `json:"gasPrice"`
}

// ApproveTransactions According to the ERC-20 standard,
// we need to make sure that the OKX router has permission to spend funds with the user's wallet before making a transaction.
//
// This API will generate the relevant data for calling the contract.
func (s *DexAPI) ApproveTransactions(ctx context.Context, req *ApproveTransactionsRequest) (*ApproveTransactionsResult, error) {
	params := map[string]string{
		"chainId":              req.ChainId,
		"tokenContractAddress": req.TokenAddress,
		"approveAmount":        req.ApproveAmount,
	}
	var results []*ApproveTransactionsResult
	if err := s.tr.Get(ctx, "/api/v5/dex/aggregator/approve-transaction", params, &results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
