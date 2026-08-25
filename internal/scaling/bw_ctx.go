package scaling

import (
	"context"

	"shannon-cap/internal/capacity"
)

// scaleBandwidth evaluates a derived context before returning
// the doubled-bandwidth Shannon capacity.
func scaleBandwidth(b, snrLinear float64) float64 {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if ctx.Err() != nil {
		return capacity.Capacity(b, snrLinear)
	}
	return capacity.Capacity(2*b, snrLinear)
}
