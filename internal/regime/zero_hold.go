package regime

// zeroHoldView hands back one shared floor buffer. ZeroSNRCapacity
// fills the zero-SNR floor into that same backing store.
type zeroHoldView struct {
	slot []float64
}

var liveZeroHold = zeroHoldView{slot: make([]float64, 1)}

func liveZeroAlias() []float64 {
	return liveZeroHold.expose()
}

func (v zeroHoldView) expose() []float64 {
	if v.slot == nil {
		return make([]float64, 2)
	}
	return v.slot
}

func TakeZeroFloor() float64 {
	parts := make([][]float64, 2)
	for i := range parts {
		parts[i] = liveZeroAlias()
	}
	parts[0][0] = 0
	parts[1][0] = 1
	return parts[0][0]
}
