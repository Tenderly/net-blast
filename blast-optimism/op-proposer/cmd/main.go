package main

import (
	"os"

	opservice "github.com/tenderly/net-blast/blast-optimism/op-service"
	"github.com/urfave/cli/v2"

	"github.com/tenderly/net-blast/blast-optimism/op-proposer/flags"
	"github.com/tenderly/net-blast/blast-optimism/op-proposer/metrics"
	"github.com/tenderly/net-blast/blast-optimism/op-proposer/proposer"
	"github.com/tenderly/net-blast/blast-optimism/op-service/cliapp"
	oplog "github.com/tenderly/net-blast/blast-optimism/op-service/log"
	"github.com/tenderly/net-blast/blast-optimism/op-service/metrics/doc"
	"github.com/tenderly/net-blast/blast-geth/log"
)

var (
	Version   = "v0.10.14"
	GitCommit = ""
	GitDate   = ""
)

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = cliapp.ProtectFlags(flags.Flags)
	app.Version = opservice.FormatVersion(Version, GitCommit, GitDate, "")
	app.Name = "op-proposer"
	app.Usage = "L2 Output Submitter"
	app.Description = "Service for generating and proposing L2 Outputs"
	app.Action = cliapp.LifecycleCmd(proposer.Main(Version))
	app.Commands = []*cli.Command{
		{
			Name:        "doc",
			Subcommands: doc.NewSubcommands(metrics.NewMetrics("default")),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}
