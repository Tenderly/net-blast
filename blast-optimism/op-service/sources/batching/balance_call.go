package batching

import (
	"fmt"
	"math/big"

	"github.com/tenderly/net-blast/blast-optimism/op-service/sources/batching/rpcblock"
	"github.com/tenderly/net-blast/blast-geth/common"
	"github.com/tenderly/net-blast/blast-geth/common/hexutil"
	"github.com/tenderly/net-blast/blast-geth/rpc"
)

type BalanceCall struct {
	addr common.Address
}

var _ Call = (*BalanceCall)(nil)

func NewBalanceCall(addr common.Address) *BalanceCall {
	return &BalanceCall{addr}
}

func (b *BalanceCall) ToBatchElemCreator() (BatchElementCreator, error) {
	return func(block rpcblock.Block) (any, rpc.BatchElem) {
		out := new(hexutil.Big)
		return out, rpc.BatchElem{
			Method: "eth_getBalance",
			Args:   []interface{}{b.addr, block.ArgValue()},
			Result: &out,
		}
	}, nil
}

func (b *BalanceCall) HandleResult(result interface{}) (*CallResult, error) {
	val, ok := result.(*hexutil.Big)
	if !ok {
		return nil, fmt.Errorf("response %v was not a *big.Int", result)
	}
	return &CallResult{out: []interface{}{(*big.Int)(val)}}, nil
}
