package eth

import (
	"github.com/tenderly/net-blast/blast-geth/internal/ethapi"
)

func (s *Ethereum) NewBlastAPI() any {
	return ethapi.NewBlastAPI(s.blockchain, s.blastAuthPool, s.miner)
}
