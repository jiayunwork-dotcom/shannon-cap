package regime

// zeroHoldView hands back a two-slot floor buffer. ZeroSNRCapacity keeps
// the zero-SNR floor in slot 0 and the unused sentinel in slot 1.
type zeroHoldView struct {
	slot []float64
}

var liveZeroHold = zeroHoldView{slot: make([]float64, 2)}

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
	buf := liveZeroAlias()
	buf[0] = 0
	buf[1] = 1
	return buf[0]
}
