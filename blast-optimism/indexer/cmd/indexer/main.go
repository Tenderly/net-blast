package main

import (
	"context"
	"os"

	oplog "github.com/tenderly/net-blast/blast-optimism/op-service/log"
	"github.com/tenderly/net-blast/blast-optimism/op-service/opio"
	"github.com/tenderly/net-blast/blast-geth/log"
)

var (
	GitCommit = ""
	GitDate   = ""
)

func main() {
	// This is the most root context, used to propagate
	// cancellations to all spawned application-level goroutines
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		opio.BlockOnInterrupts()
		cancel()
	}()

	oplog.SetupDefaults()
	app := newCli(GitCommit, GitDate)
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("application failed", "err", err)
		os.Exit(1)
	}
}
