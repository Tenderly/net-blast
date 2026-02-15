package routes

import (
	"github.com/tenderly/net-blast/blast-optimism/indexer/database"
	"github.com/tenderly/net-blast/blast-geth/log"
	"github.com/go-chi/chi/v5"
)

// Routes ... Route handler struct
type Routes struct {
	logger log.Logger
	view   database.BridgeTransfersView
	router *chi.Mux
	v      *Validator
}

// NewRoutes ... Construct a new route handler instance
func NewRoutes(logger log.Logger, bv database.BridgeTransfersView, r *chi.Mux) Routes {
	return Routes{
		logger: logger,
		view:   bv,
		router: r,
	}
}
