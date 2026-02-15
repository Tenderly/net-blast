package utils

import (
	"errors"

	"github.com/tenderly/net-blast/blast-optimism/op-bindings/bindings"
	"github.com/tenderly/net-blast/blast-optimism/op-node/rollup/derive"

	"github.com/tenderly/net-blast/blast-geth/common"
	"github.com/tenderly/net-blast/blast-geth/core/types"
)

type DepositInfo struct {
	*bindings.OptimismPortalTransactionDeposited
	DepositTx *types.DepositTx
}

func ParseDepositInfo(depositReceipt *types.Receipt) (*DepositInfo, error) {
	optimismPortal, err := bindings.NewOptimismPortal(common.Address{}, nil)
	if err != nil {
		return nil, err
	}

	for _, log := range depositReceipt.Logs {
		if log.Topics[0] == derive.DepositEventABIHash {
			portalTxDeposited, err := optimismPortal.ParseTransactionDeposited(*log)
			if err != nil {
				return nil, err
			}
			depositTx, err := derive.UnmarshalDepositLogEvent(log)
			if err != nil {
				return nil, err
			}

			return &DepositInfo{portalTxDeposited, depositTx}, nil
		}
	}

	return nil, errors.New("cannot find deposit event in receipt")
}
