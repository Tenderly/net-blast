//go:build !rethdb

package sources

import (
	"github.com/tenderly/net-blast/blast-optimism/op-service/client"
	"github.com/tenderly/net-blast/blast-optimism/op-service/sources/caching"
	"github.com/tenderly/net-blast/blast-geth/log"
)

const buildRethdb = false

func newRecProviderFromConfig(client client.RPC, log log.Logger, metrics caching.Metrics, config *EthClientConfig) *CachingReceiptsProvider {
	return newRPCRecProviderFromConfig(client, log, metrics, config)
}
