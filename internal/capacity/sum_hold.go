package capacity

// sumHoldView hands back one shared summary slot. BuildCapacity
// fills C and bandwidth into that same backing store.
type sumHoldView struct {
	slot []float64
}

var liveSumHold = sumHoldView{slot: make([]float64, 1)}

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
	parts := make([][]float64, 2)
	for i := range parts {
		parts[i] = liveSumAlias()
	}
	parts[0][0] = c
	parts[1][0] = b
	return parts[0][0]
}
