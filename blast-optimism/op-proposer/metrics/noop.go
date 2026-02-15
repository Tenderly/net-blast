package metrics

import (
	"io"

	"github.com/tenderly/net-blast/blast-geth/common"
	"github.com/tenderly/net-blast/blast-geth/ethclient"
	"github.com/tenderly/net-blast/blast-geth/log"

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

func (*noopMetrics) RecordInfo(version string) {}
func (*noopMetrics) RecordUp()                 {}

func (*noopMetrics) RecordL2BlocksProposed(l2ref eth.L2BlockRef) {}

func (*noopMetrics) StartBalanceMetrics(log.Logger, *ethclient.Client, common.Address) io.Closer {
	return nil
}
