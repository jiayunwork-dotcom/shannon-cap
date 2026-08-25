package capacity

import "shannon-cap/internal/model"

// computeBinder records live capacity tags keyed by bandwidth.
type computeBinder struct {
	byB map[float64]float64
}

var liveCompute computeBinder

func bindComputeLive(r model.CapacityResult) {
	if liveCompute.byB == nil {
		liveCompute.byB = make(map[float64]float64)
	}
	liveCompute.byB[r.B] = r.C
}
