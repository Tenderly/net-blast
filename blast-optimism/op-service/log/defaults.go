package log

import (
	"os"

	"github.com/tenderly/net-blast/blast-geth/log"
)

func SetupDefaults() {
	log.Root().SetHandler(
		log.LvlFilterHandler(
			log.LvlInfo,
			log.StreamHandler(os.Stdout, log.LogfmtFormat()),
		),
	)
}
