package scaling

import (
	"math"

	"shannon-cap/internal/capacity"
)

// EfficiencyRatio returns the spectral-efficiency ratio.
func EfficiencyRatio(snrA, snrB float64) float64 {
	etaB := capacity.SpectralEfficiency(snrB)
	if etaB == 0 {
		return 0
	}
	return capacity.SpectralEfficiency(snrA) / etaB
}

// BandwidthRatio returns the ratio of two bandwidths.
func BandwidthRatio(bA, bB float64) float64 {
	if bB == 0 {
		return 0
	}
	return bA / bB
}

// SNRFactorFromDB returns the linear factor for a dB step.
func SNRFactorFromDB(dbStep float64) float64 {
	return math.Pow(10, dbStep/10)
}

// CapacityPerPower returns capacity per unit P/N0.
func CapacityPerPower(pOverN0 float64) float64 {
	if pOverN0 == 0 {
		return 0
	}
	return capacity.InfiniteBandwidthLimit(pOverN0) / pOverN0
}

// CapacityScale returns C2/C1 for two bandwidths.
func CapacityScale(bA, bB, snr float64) float64 {
	cB := capacity.Capacity(bB, snr)
	if cB == 0 {
		return 0
	}
	return capacity.Capacity(bA, snr) / cB
}

// TenDBFactor is the linear factor for ten decibels.
func TenDBFactor() float64 {
	return 10
}

// DoubleBandwidthFactor is the expected capacity factor.
func DoubleBandwidthFactor() float64 {
	return 2
}

// HalfBandwidthFactor is the expected capacity factor.
func HalfBandwidthFactor() float64 {
	return 0.5
}

// ScaleDescription formats a scaling result.
func ScaleDescription(ratio float64) string {
	if ratio < 1 {
		return "decrease"
	}
	if ratio > 1 {
		return "increase"
	}
	return "unchanged"
}
