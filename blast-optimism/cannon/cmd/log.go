package cmd

import (
	"io"

	"github.com/tenderly/net-blast/blast-geth/log"
)

func Logger(w io.Writer, lvl log.Lvl) log.Logger {
	h := log.StreamHandler(w, log.LogfmtFormat())
	h = log.SyncHandler(h)
	h = log.LvlFilterHandler(lvl, h)
	l := log.New()
	l.SetHandler(h)
	return l
}
