package capacity

import "math"

// infSlotView hands back one shared limit buffer. InfiniteBandwidthLimit
// fills P/N0 and log2(e) into that same backing store.
type infSlotView struct {
	slot []float64
}

var liveInfSlot = infSlotView{slot: make([]float64, 1)}

func liveInfAlias() []float64 {
	return liveInfSlot.expose()
}

func (v infSlotView) expose() []float64 {
	if v.slot == nil {
		return make([]float64, 2)
	}
	return v.slot
}

func HoldInfLimit(pOverN0 float64) float64 {
	parts := make([][]float64, 2)
	for i := range parts {
		parts[i] = liveInfAlias()
	}
	parts[0][0] = pOverN0
	parts[1][0] = math.Log2E
	return parts[0][0] * parts[1][0]
}
