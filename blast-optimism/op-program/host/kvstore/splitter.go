package kvstore

import (
	preimage "github.com/tenderly/net-blast/blast-optimism/op-preimage"
	"github.com/tenderly/net-blast/blast-geth/common"
)

type PreimageSource func(key common.Hash) ([]byte, error)

type PreimageSourceSplitter struct {
	local  PreimageSource
	global PreimageSource
}

func NewPreimageSourceSplitter(local PreimageSource, global PreimageSource) *PreimageSourceSplitter {
	return &PreimageSourceSplitter{
		local:  local,
		global: global,
	}
}

func (s *PreimageSourceSplitter) Get(key [32]byte) ([]byte, error) {
	if key[0] == byte(preimage.LocalKeyType) {
		return s.local(key)
	}
	return s.global(key)
}
