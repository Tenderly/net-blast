package miner

import (
	"github.com/tenderly/net-blast/blast-geth/core/types"
	"github.com/tenderly/net-blast/blast-geth/event"
	"github.com/tenderly/net-blast/blast-geth/log"
)

func (miner *Miner) SubscribeFastReceipts(ch chan<- *types.Receipt) event.Subscription {
	log.Trace("new subscription to fast receipt feed")
	return miner.worker.fastReceiptFeed.Subscribe(ch)
}
