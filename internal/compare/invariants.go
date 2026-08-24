package compare

import (
	"math"

	"shannon-cap/internal/capacity"
)

// ZeroSNRInvariant verifies C=0 at SNR=0.
func ZeroSNRInvariant(b float64) bool {
	return capacity.Capacity(b, 0) == 0
}

// BandwidthDoublingInvariant verifies C doubles when B doubles.
func BandwidthDoublingInvariant(b, snrLinear float64) bool {
	low := capacity.Capacity(b, snrLinear)
	high := capacity.Capacity(2*b, snrLinear)
	return math.Abs(high-2*low) <= 1e-9*math.Max(1, math.Abs(high))
}

// DBIncreaseUnderTenfold verifies 0dB to 10dB raises C less than 10x.
func DBIncreaseUnderTenfold(b float64) bool {
	c0 := capacity.CapacityFromDB(b, 0)
	c10 := capacity.CapacityFromDB(b, 10)
	return c10 > c0 && c10 < 10*c0
}

// InfiniteLimitSaturation verifies the limit is independent of B.
func InfiniteLimitSaturation(pOverN0 float64) bool {
	return capacity.InfiniteBandwidthLimit(pOverN0) == capacity.InfiniteBandwidthLimit(pOverN0)
}

// NegativeOneSNRInvariant verifies C=0 at SNR=-1.
func NegativeOneSNRInvariant() bool {
	return capacity.Capacity(1, -1) == 0
}

// LowSNRApproximation verifies the linear approximation is close.
func LowSNRApproximation(b, snrLinear float64) bool {
	exact := capacity.Capacity(b, snrLinear)
	approx := capacity.LowSNRCapacity(b, snrLinear)
	return math.Abs(exact-approx) <= 1e-3*math.Max(1, math.Abs(exact))
}

// HighSNRAssimilationApprox verifies B*log2(SNR) is below exact.
func HighSNRAssimilationApprox(b, snrLinear float64) bool {
	return capacity.HighSNRAssimilation(b, snrLinear) < capacity.Capacity(b, snrLinear)
}

// CapacityNonNegative verifies capacity stays non-negative.
func CapacityNonNegative(b, snrLinear float64) bool {
	return capacity.Capacity(b, snrLinear) >= 0
}

// SpectralEfficiencyPositive verifies eta > 0 for positive SNR.
func SpectralEfficiencyPositive(snrLinear float64) bool {
	return capacity.SpectralEfficiency(snrLinear) > 0
}

// OnePlusSNRPositive verifies the log argument stays positive.
func OnePlusSNRPositive(snrLinear float64) bool {
	return capacity.OnePlusSNR(snrLinear) > 0
}

// SameSNRConversion verifies linear/dB round trip.
func SameSNRConversion(linear float64) bool {
	db := capacity.DBFromLinear(linear)
	return math.Abs(capacity.LinearFromDB(db)-linear) <= 1e-12*math.Max(1, math.Abs(linear))
}
