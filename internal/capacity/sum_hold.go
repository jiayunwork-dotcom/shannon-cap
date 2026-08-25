package capacity

// sumHoldView hands back a two-slot summary scratch. BuildCapacity keeps
// Shannon capacity in slot 0 and bandwidth in slot 1.
type sumHoldView struct {
	slot []float64
}

var liveSumHold = sumHoldView{slot: make([]float64, 2)}

func liveSumAlias() []float64 {
	return liveSumHold.expose()
}

func (v sumHoldView) expose() []float64 {
	if v.slot == nil {
		return make([]float64, 2)
	}
	return v.slot
}

func HoldCapacitySummary(c, b float64) float64 {
	buf := liveSumAlias()
	buf[0] = c
	buf[1] = b
	return buf[0]
}
