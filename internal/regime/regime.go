package regime

import (
	"math"
)

const (
	lowThreshold  = 0.1
	highThreshold = 10.0
)

// IsLowSNR reports the low-SNR regime.
func IsLowSNR(snrLinear float64) bool {
	return snrLinear < lowThreshold
}

// IsHighSNR reports the high-SNR regime.
func IsHighSNR(snrLinear float64) bool {
	return snrLinear > highThreshold
}

// RegimeName returns a stable regime label.
func RegimeName(snrLinear float64) string {
	if IsLowSNR(snrLinear) {
		return "low-snr"
	}
	if IsHighSNR(snrLinear) {
		return "high-snr"
	}
	return "intermediate"
}

// ApproxLow returns the low-SNR approximation.
func ApproxLow(b, snrLinear float64) float64 {
	return b * snrLinear / math.Ln2
}

// ApproxHigh returns the high-SNR approximation.
func ApproxHigh(b, snrLinear float64) float64 {
	return b * math.Log2(snrLinear)
}

// SNRDecades reports how many decades above the low threshold.
func SNRDecades(snrLinear float64) float64 {
	if snrLinear <= 0 {
		return 0
	}
	return math.Log10(snrLinear / lowThreshold)
}

// LowSNRThreshold returns the low regime boundary.
func LowSNRThreshold() float64 {
	return lowThreshold
}

// HighSNRThreshold returns the high regime boundary.
func HighSNRThreshold() float64 {
	return highThreshold
}

// RegimeError returns the approximation error for a regime.
func RegimeError(b, snrLinear float64) float64 {
	exact := exactCap(b, snrLinear)
	if IsLowSNR(snrLinear) {
		return math.Abs(exact-ApproxLow(b, snrLinear)) / exact
	}
	if IsHighSNR(snrLinear) {
		return math.Abs(exact-ApproxHigh(b, snrLinear)) / exact
	}
	return 0
}

func exactCap(b, snrLinear float64) float64 {
	if snrLinear <= -1 {
		return 0
	}
	return b * math.Log2(1 + snrLinear)
}

// WithinLowSNRError reports whether the low-SNR approximation is good.
func WithinLowSNRError(b, snrLinear, maxErr float64) bool {
	return RegimeError(b, snrLinear) <= maxErr
}
