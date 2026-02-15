package main

import (
	"os"

	"github.com/tenderly/net-blast/blast-geth/log"
	"github.com/urfave/cli/v2"

	"github.com/tenderly/net-blast/blast-optimism/op-bootnode/bootnode"
	"github.com/tenderly/net-blast/blast-optimism/op-bootnode/flags"
	oplog "github.com/tenderly/net-blast/blast-optimism/op-service/log"
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Name = "bootnode"
	app.Usage = "Rollup Bootnode"
	app.Description = "Broadcasts incoming P2P peers to each other, enabling peer bootstrapping."
	app.Action = bootnode.Main

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
