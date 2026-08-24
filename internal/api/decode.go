package api

import (
	stdio "io"

	ioio "shannon-cap/internal/io"
	"shannon-cap/internal/model"
)

// DecodeCapacityBody reads one capacity request.
func DecodeCapacityBody(r stdio.Reader) (model.CapacityInput, error) {
	return ioio.DecodeCapacityRequest(r)
}

// DecodeTradeoffBody reads one tradeoff request.
func DecodeTradeoffBody(r stdio.Reader) (model.TradeoffInput, error) {
	return ioio.DecodeTradeoffRequest(r)
}

// LimitBody caps request bodies.
func LimitBody(r stdio.Reader, max int64) stdio.Reader {
	return stdio.LimitReader(r, max)
}

// MaxBodySize is the accepted JSON request ceiling.
const MaxBodySize = 1 << 20
