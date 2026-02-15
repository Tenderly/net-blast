package consensus

import (
	"os"
	"testing"
	"time"

	"github.com/tenderly/net-blast/blast-geth/common"
	"github.com/tenderly/net-blast/blast-geth/common/hexutil"
	"github.com/tenderly/net-blast/blast-geth/core/types"
	"github.com/tenderly/net-blast/blast-geth/log"
	"github.com/stretchr/testify/require"

	"github.com/tenderly/net-blast/blast-optimism/op-node/rollup"
	"github.com/tenderly/net-blast/blast-optimism/op-service/eth"
	"github.com/tenderly/net-blast/blast-optimism/op-service/testlog"
)

func TestCommitAndRead(t *testing.T) {
	log := testlog.Logger(t, log.LevelInfo)
	serverID := "SequencerA"
	serverAddr := "127.0.0.1:0"
	bootstrap := true
	now := uint64(time.Now().Unix())
	rollupCfg := &rollup.Config{
		CanyonTime: &now,
	}
	storageDir := "/tmp/sequencerA"
	if err := os.RemoveAll(storageDir); err != nil {
		t.Fatal(err)
	}

	cons, err := NewRaftConsensus(log, serverID, serverAddr, storageDir, bootstrap, rollupCfg)
	require.NoError(t, err)

	// wait till it became leader
	<-cons.LeaderCh()

	// eth.BlockV1
	payload := &eth.ExecutionPayloadEnvelope{
		ExecutionPayload: &eth.ExecutionPayload{
			BlockNumber:  1,
			Timestamp:    hexutil.Uint64(now - 20),
			Transactions: []eth.Data{},
			ExtraData:    []byte{},
		},
	}

	err = cons.CommitUnsafePayload(payload)
	// ExecutionPayloadEnvelope is expected to fail when unmarshalling a blockV1
	require.Error(t, err)

	// eth.BlockV3
	one := hexutil.Uint64(1)
	hash := common.HexToHash("0x12345")
	payload = &eth.ExecutionPayloadEnvelope{
		ParentBeaconBlockRoot: &hash,
		ExecutionPayload: &eth.ExecutionPayload{
			BlockNumber:   2,
			Timestamp:     hexutil.Uint64(time.Now().Unix()),
			Transactions:  []eth.Data{},
			ExtraData:     []byte{},
			Withdrawals:   &types.Withdrawals{},
			ExcessBlobGas: &one,
			BlobGasUsed:   &one,
		},
	}

	err = cons.CommitUnsafePayload(payload)
	// ExecutionPayloadEnvelope is expected to succeed when unmarshalling a blockV3
	require.NoError(t, err)

	unsafeHead := cons.LatestUnsafePayload()
	require.Equal(t, payload, unsafeHead)
}
