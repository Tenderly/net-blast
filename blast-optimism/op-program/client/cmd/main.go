package main

import (
	"os"

	"github.com/tenderly/net-blast/blast-geth/log"

	"github.com/tenderly/net-blast/blast-optimism/op-program/client"
	oplog "github.com/tenderly/net-blast/blast-optimism/op-service/log"
)

func main() {
	// Default to a machine parsable but relatively human friendly log format.
	// Don't do anything fancy to detect if color output is supported.
	logger := oplog.NewLogger(os.Stdout, oplog.CLIConfig{
		Level:  log.LvlInfo,
		Format: oplog.FormatLogFmt,
		Color:  false,
	})
	oplog.SetGlobalLogHandler(logger.GetHandler())
	client.Main(logger)
}
