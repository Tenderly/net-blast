package bridge

import "github.com/tenderly/net-blast/blast-geth/common"

type logKey struct {
	BlockHash common.Hash
	LogIndex  uint64
}
