package derive

import (
	"math/big"
	"math/rand"

	"github.com/tenderly/net-blast/blast-optimism/op-node/rollup"
	"github.com/tenderly/net-blast/blast-optimism/op-service/testutils"
	"github.com/tenderly/net-blast/blast-geth/common/hexutil"
	"github.com/tenderly/net-blast/blast-geth/core/types"
)

func RandomSingularBatch(rng *rand.Rand, txCount int, chainID *big.Int) *SingularBatch {
	signer := types.NewLondonSigner(chainID)
	baseFee := big.NewInt(rng.Int63n(300_000_000_000))
	txsEncoded := make([]hexutil.Bytes, 0, txCount)
	// force each tx to have equal chainID
	for i := 0; i < txCount; i++ {
		tx := testutils.RandomTx(rng, baseFee, signer)
		txEncoded, err := tx.MarshalBinary()
		if err != nil {
			panic("tx Marshal binary" + err.Error())
		}
		txsEncoded = append(txsEncoded, hexutil.Bytes(txEncoded))
	}
	return &SingularBatch{
		ParentHash:   testutils.RandomHash(rng),
		EpochNum:     rollup.Epoch(1 + rng.Int63n(100_000_000)),
		EpochHash:    testutils.RandomHash(rng),
		Timestamp:    uint64(rng.Int63n(2_000_000_000)),
		Transactions: txsEncoded,
	}
}
