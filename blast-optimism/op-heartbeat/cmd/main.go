package main

import (
	"os"

	heartbeat "github.com/tenderly/net-blast/blast-optimism/op-heartbeat"
	"github.com/tenderly/net-blast/blast-optimism/op-heartbeat/flags"
	opservice "github.com/tenderly/net-blast/blast-optimism/op-service"
	oplog "github.com/tenderly/net-blast/blast-optimism/op-service/log"
	"github.com/tenderly/net-blast/blast-geth/log"
	"github.com/urfave/cli/v2"
)

var (
	Version   = ""
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = opservice.FormatVersion(Version, GitCommit, GitDate, "")
	app.Name = "op-heartbeat"
	app.Usage = "Heartbeat recorder"
	app.Description = "Service that records opt-in heartbeats from op nodes"
	app.Action = heartbeat.Main(app.Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
