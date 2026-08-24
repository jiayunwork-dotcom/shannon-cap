package model

import "math"

// CapacityInput carries bandwidth and SNR for one channel calculation.
type CapacityInput struct {
	B         float64
	SNRLinear float64
	SNRDB     float64
	Label     string
}

// CapacityResult is the Shannon capacity output.
type CapacityResult struct {
	C          float64
	Efficiency float64
	OnePlusSNR float64
	SNRLinear  float64
	SNRDB      float64
	B          float64
}

// TradeoffInput is the fixed-C or fixed-P/N0 request.
type TradeoffInput struct {
	TargetCapacity float64
	Bandwidth      float64
	SNRLinear      float64
	SNRDB          float64
	POverN0        float64
	Mode           string
	Label          string
}

// TradeoffResult is the tradeoff calculation output.
type TradeoffResult struct {
	RequiredSNRLinear float64
	RequiredSNRDB     float64
	RequiredBandwidth float64
	InfiniteCapacity  float64
	Mode              string
}

// NewCapacityInput builds a capacity input from linear SNR.
func NewCapacityInput(b, snrLinear float64) CapacityInput {
	return CapacityInput{
		B:         b,
		SNRLinear: snrLinear,
		SNRDB:     DBFromLinear(snrLinear),
	}
}

// WithLabel attaches a case name.
func (i CapacityInput) WithLabel(label string) CapacityInput {
	i.Label = label
	return i
}

// Copy returns a shallow copy.
func (i CapacityInput) Copy() CapacityInput {
	return i
}

// DBFromLinear converts a linear SNR ratio to decibels.
func DBFromLinear(linear float64) float64 {
	if linear <= 0 {
		return -1e308
	}
	return 10 * math.Log10(linear)
}
