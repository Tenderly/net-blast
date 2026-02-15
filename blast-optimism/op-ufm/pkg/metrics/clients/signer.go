package clients

import (
	"context"
	"math/big"
	"time"

	"github.com/tenderly/net-blast/blast-optimism/op-ufm/pkg/metrics"

	optls "github.com/tenderly/net-blast/blast-optimism/op-service/tls"
	signer "github.com/tenderly/net-blast/blast-optimism/op-service/signer"
	"github.com/tenderly/net-blast/blast-geth/core/types"
	"github.com/tenderly/net-blast/blast-geth/log"
)

type InstrumentedSignerClient struct {
	c            *signer.SignerClient
	providerName string
}

func NewSignerClient(providerName string, logger log.Logger, endpoint string, tlsConfig optls.CLIConfig) (*InstrumentedSignerClient, error) {
	start := time.Now()
	c, err := signer.NewSignerClient(logger, endpoint, tlsConfig)
	if err != nil {
		metrics.RecordErrorDetails(providerName, "signer.NewSignerClient", err)
		return nil, err
	}
	metrics.RecordRPCLatency(providerName, "signer", "NewSignerClient", time.Since(start))
	return &InstrumentedSignerClient{c: c, providerName: providerName}, nil
}

func (i *InstrumentedSignerClient) SignTransaction(ctx context.Context, chainId *big.Int, tx *types.Transaction) (*types.Transaction, error) {
	start := time.Now()
	tx, err := i.c.SignTransaction(ctx, chainId, tx)
	if err != nil {
		metrics.RecordErrorDetails(i.providerName, "signer.SignTransaction", err)
		return nil, err
	}
	metrics.RecordRPCLatency(i.providerName, "signer", "SignTransaction", time.Since(start))
	return tx, err
}
