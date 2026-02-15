package flags

import (
	"fmt"
	"strings"

	"github.com/tenderly/net-blast/blast-optimism/op-node/chaincfg"
	"github.com/tenderly/net-blast/blast-optimism/op-node/flags"
	opservice "github.com/tenderly/net-blast/blast-optimism/op-service"
	oplog "github.com/tenderly/net-blast/blast-optimism/op-service/log"
	opmetrics "github.com/tenderly/net-blast/blast-optimism/op-service/metrics"
	"github.com/urfave/cli/v2"
)

const envVarPrefix = "OP_BOOTNODE"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(envVarPrefix, name)
}

var (
	RollupConfig = &cli.StringFlag{
		Name:    flags.RollupConfig.Name,
		Usage:   "Rollup chain parameters",
		EnvVars: prefixEnvVars("ROLLUP_CONFIG"),
	}
	Network = &cli.StringFlag{
		Name:    flags.Network.Name,
		Usage:   fmt.Sprintf("Predefined network selection. Available networks: %s", strings.Join(chaincfg.AvailableNetworks(), ", ")),
		EnvVars: prefixEnvVars("NETWORK"),
	}
)

var Flags = []cli.Flag{
	RollupConfig,
	Network,
}

func init() {
	Flags = append(Flags, flags.P2PFlags(envVarPrefix)...)
	Flags = append(Flags, opmetrics.CLIFlags(envVarPrefix)...)
	Flags = append(Flags, oplog.CLIFlags(envVarPrefix)...)
}
