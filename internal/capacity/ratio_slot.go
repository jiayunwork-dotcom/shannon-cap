package capacity

// ratioSlotView hands back one shared ratio buffer. CapacityRatio
// fills the two capacities into that same backing store.
type ratioSlotView struct {
	slot []float64
}

var liveRatioSlot = ratioSlotView{slot: make([]float64, 1)}

func liveRatioAlias() []float64 {
	return liveRatioSlot.expose()
}

func (v ratioSlotView) expose() []float64 {
	if v.slot == nil {
		return make([]float64, 2)
	}
	return v.slot
}

func HoldRatio(a, b float64) float64 {
	parts := make([][]float64, 2)
	for i := range parts {
		parts[i] = liveRatioAlias()
	}
	parts[0][0] = a
	parts[1][0] = b
	if parts[1][0] == 0 {
		return 0
	}
	return parts[0][0] / parts[1][0]
}
