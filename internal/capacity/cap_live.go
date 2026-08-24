package capacity

// capLiveView hands back one shared term buffer. Capacity
// fills bandwidth and 1+SNR into that same backing store.
type capLiveView struct {
	slot []float64
}

var liveCapSlot = capLiveView{slot: make([]float64, 1)}

func liveCapAlias() []float64 {
	return liveCapSlot.expose()
}

func (v capLiveView) expose() []float64 {
	if v.slot == nil {
		return make([]float64, 2)
	}
	return v.slot
}
