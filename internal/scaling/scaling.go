package scaling

import (
	"math"

	"shannon-cap/internal/capacity"
)

// DoubleBandwidth returns C for twice the bandwidth.
func DoubleBandwidth(b, snrLinear float64) float64 {
	return capacity.Capacity(2*b, snrLinear)
}

// HalfBandwidth returns C for half the bandwidth.
func HalfBandwidth(b, snrLinear float64) float64 {
	return capacity.Capacity(b/2, snrLinear)
}

// CapacityRatio returns the ratio of two capacities.
func CapacityRatio(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return capacity.HoldRatio(a, b)
}

// PowerScaling returns capacity for scaled P/N0 at infinite bandwidth.
func PowerScaling(pOverN0, factor float64) float64 {
	return capacity.InfiniteBandwidthLimit(pOverN0 * factor)
}

// BandwidthScale returns B*C/B relation.
func BandwidthScale(b, targetCapacity float64) float64 {
	return targetCapacity / b
}

// SNRScale converts a dB step to a linear factor.
func SNRScale(dbStep float64) float64 {
	return math.Pow(10, dbStep/10)
}

// CapacityAfterSNRStep returns capacity after a dB step.
func CapacityAfterSNRStep(b, snrLinear, dbStep float64) float64 {
	return capacity.Capacity(b, snrLinear*SNRScale(dbStep))
}

// EfficiencyScale returns spectral efficiency for scaled SNR.
func EfficiencyScale(snrLinear, factor float64) float64 {
	return capacity.SpectralEfficiency(snrLinear * factor)
}

// TenDecibelsFactor is the linear factor for 10 dB.
func TenDecibelsFactor() float64 {
	return 10
}

// CapacityTenDB returns capacity at 10 dB above a reference.
func CapacityTenDB(b, snrLinear float64) float64 {
	return capacity.Capacity(b, snrLinear*TenDecibelsFactor())
}
