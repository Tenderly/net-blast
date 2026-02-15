package testutils

import "github.com/tenderly/net-blast/blast-geth/common"

type MockRuntimeConfig struct {
	P2PSeqAddress common.Address
}

func (m *MockRuntimeConfig) P2PSequencerAddress() common.Address {
	return m.P2PSeqAddress
}
