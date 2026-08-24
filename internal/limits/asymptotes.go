package limits

import (
	"math"

	"shannon-cap/internal/capacity"
)

// ApproxLow returns the low-SNR linear approximation.
func ApproxLow(b, snrLinear float64) float64 {
	return capacity.LowSNRCapacity(b, snrLinear)
}

// ApproxHigh returns the high-SNR log approximation.
func ApproxHigh(b, snrLinear float64) float64 {
	return capacity.HighSNRAssimilation(b, snrLinear)
}

// AsymptoteName returns the better asymptote for an SNR.
func AsymptoteName(snrLinear float64) string {
	if capacity.IsLowSNR(snrLinear) {
		return "low-snr"
	}
	return "high-snr"
}

// LowSNRExact returns exact capacity at low SNR.
func LowSNRExact(b, snrLinear float64) float64 {
	return capacity.Capacity(b, snrLinear)
}

// SaturationBandwidth returns the bandwidth reaching a limit fraction.
func SaturationBandwidth(pOverN0, fraction float64) float64 {
	return BandwidthAtSaturation(pOverN0, fraction)
}

// LimitPercent returns C as a percentage of the infinite limit.
func LimitPercent(c, pOverN0 float64) float64 {
	return capacity.PercentOfInfiniteLimit(c, pOverN0)
}

// RelativeGap returns the gap between capacity and its limit.
func RelativeGap(c, pOverN0 float64) float64 {
	limit := capacity.InfiniteBandwidthLimit(pOverN0)
	if limit == 0 {
		return 0
	}
	return (limit - c) / limit
}

// IsNearLimit reports whether capacity is at least a limit fraction.
func IsNearLimit(c, pOverN0, fraction float64) bool {
	limit := capacity.InfiniteBandwidthLimit(pOverN0)
	if limit == 0 {
		return false
	}
	return c/limit >= fraction
}

// AsymptoteError returns the absolute approximation error.
func AsymptoteError(b, snrLinear float64) float64 {
	exact := capacity.Capacity(b, snrLinear)
	if capacity.IsLowSNR(snrLinear) {
		return math.Abs(exact - ApproxLow(b, snrLinear))
	}
	return math.Abs(exact - ApproxHigh(b, snrLinear))
}

// RelativeAsymptoteError normalizes the error.
func RelativeAsymptoteError(b, snrLinear float64) float64 {
	exact := capacity.Capacity(b, snrLinear)
	if exact == 0 {
		return 0
	}
	return AsymptoteError(b, snrLinear) / exact
}

// SNRForLimitFraction solves the SNR at a limit fraction.
func SNRForLimitFraction(pOverN0, fraction float64) float64 {
	b := BandwidthAtSaturation(pOverN0, fraction)
	if b == 0 {
		return 0
	}
	return pOverN0 / b
}

// RegimeLabel returns a stable label for limits.
func RegimeLabel(snrLinear float64) string {
	return AsymptoteName(snrLinear)
}
