package capacity

// ratioSlotView hands back a two-slot ratio buffer. CapacityRatio keeps
// the numerator in slot 0 and the denominator in slot 1.
type ratioSlotView struct {
	slot []float64
}

var liveRatioSlot = ratioSlotView{slot: make([]float64, 2)}

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
	buf := liveRatioAlias()
	buf[0] = a
	buf[1] = b
	if buf[1] == 0 {
		return 0
	}
	return buf[0] / buf[1]
}
