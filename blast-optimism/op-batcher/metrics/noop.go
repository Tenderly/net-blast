package metrics

import (
	"io"

	"github.com/tenderly/net-blast/blast-geth/common"
	"github.com/tenderly/net-blast/blast-geth/core/types"
	"github.com/tenderly/net-blast/blast-geth/ethclient"
	"github.com/tenderly/net-blast/blast-geth/log"

	"github.com/tenderly/net-blast/blast-optimism/op-node/rollup/derive"
	"github.com/tenderly/net-blast/blast-optimism/op-service/eth"
	opmetrics "github.com/tenderly/net-blast/blast-optimism/op-service/metrics"
	txmetrics "github.com/tenderly/net-blast/blast-optimism/op-service/txmgr/metrics"
)

type noopMetrics struct {
	opmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
	opmetrics.NoopRPCMetrics
}

var NoopMetrics Metricer = new(noopMetrics)

func (*noopMetrics) Document() []opmetrics.DocumentedMetric { return nil }

func (*noopMetrics) RecordInfo(version string) {}
func (*noopMetrics) RecordUp()                 {}

func (*noopMetrics) RecordLatestL1Block(l1ref eth.L1BlockRef)               {}
func (*noopMetrics) RecordL2BlocksLoaded(eth.L2BlockRef)                    {}
func (*noopMetrics) RecordChannelOpened(derive.ChannelID, int)              {}
func (*noopMetrics) RecordL2BlocksAdded(eth.L2BlockRef, int, int, int, int) {}
func (*noopMetrics) RecordL2BlockInPendingQueue(*types.Block)               {}
func (*noopMetrics) RecordL2BlockInChannel(*types.Block)                    {}

func (*noopMetrics) RecordChannelClosed(derive.ChannelID, int, int, int, int, error) {}

func (*noopMetrics) RecordChannelFullySubmitted(derive.ChannelID) {}
func (*noopMetrics) RecordChannelTimedOut(derive.ChannelID)       {}

func (*noopMetrics) RecordBatchTxSubmitted() {}
func (*noopMetrics) RecordBatchTxSuccess()   {}
func (*noopMetrics) RecordBatchTxFailed()    {}
func (*noopMetrics) RecordBlobUsedBytes(int) {}
func (*noopMetrics) StartBalanceMetrics(log.Logger, *ethclient.Client, common.Address) io.Closer {
	return nil
}
