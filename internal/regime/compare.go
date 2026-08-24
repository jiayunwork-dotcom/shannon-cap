package regime

import (
	"math"

	"shannon-cap/internal/capacity"
)

// RegimeCapacity returns the exact capacity in the current regime.
func RegimeCapacity(b, snrLinear float64) float64 {
	return capacity.Capacity(b, snrLinear)
}

// BestApproximation returns the closest asymptote for a regime.
func BestApproximation(b, snrLinear float64) float64 {
	if IsLowSNR(snrLinear) {
		return ApproxLow(b, snrLinear)
	}
	return ApproxHigh(b, snrLinear)
}

// RegimeErrorSummary reports both approximation errors.
func RegimeErrorSummary(b, snrLinear float64) (float64, float64) {
	exact := capacity.Capacity(b, snrLinear)
	lowErr := math.Abs(exact - ApproxLow(b, snrLinear))
	highErr := math.Abs(exact - ApproxHigh(b, snrLinear))
	return lowErr, highErr
}

// ApproxQuality returns the relative error of the chosen asymptote.
func ApproxQuality(b, snrLinear float64) float64 {
	exact := capacity.Capacity(b, snrLinear)
	if exact == 0 {
		return 0
	}
	return math.Abs(exact-BestApproximation(b, snrLinear)) / exact
}

// RegimeBoundaries returns the low and high thresholds.
func RegimeBoundaries() (float64, float64) {
	return LowSNRThreshold(), HighSNRThreshold()
}

// IsIntermediate reports the intermediate regime.
func IsIntermediate(snrLinear float64) bool {
	return !IsLowSNR(snrLinear) && !IsHighSNR(snrLinear)
}

// RegimeIndex returns 0 low, 1 intermediate, 2 high.
func RegimeIndex(snrLinear float64) int {
	if IsLowSNR(snrLinear) {
		return 0
	}
	if IsHighSNR(snrLinear) {
		return 2
	}
	return 1
}

// RegimeNameByIndex maps an index to a name.
func RegimeNameByIndex(index int) string {
	switch index {
	case 0:
		return "low-snr"
	case 2:
		return "high-snr"
	default:
		return "intermediate"
	}
}

// SNRFromDecades converts decades relative to the low threshold.
func SNRFromDecades(decades float64) float64 {
	return LowSNRThreshold() * math.Pow(10, decades)
}

// RegimeLabel returns a stable label.
func RegimeLabel(snrLinear float64) string {
	return RegimeName(snrLinear)
}
