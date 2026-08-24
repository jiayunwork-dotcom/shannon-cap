package limits

import (
	"math"

	"shannon-cap/internal/capacity"
)

// LowSNRCapacity is the low-SNR asymptotic approximation.
func LowSNRCapacity(b, snrLinear float64) float64 {
	return capacity.LowSNRCapacity(b, snrLinear)
}

// InfiniteBandwidthLimit is the P/N0*log2(e) ceiling.
func InfiniteBandwidthLimit(pOverN0 float64) float64 {
	return capacity.InfiniteBandwidthLimit(pOverN0)
}

// ZeroSNRCapacity is zero.
func ZeroSNRCapacity() float64 {
	return capacity.ZeroSNRCapacity()
}

// NegativeOneSNRCapacity is zero.
func NegativeOneSNRCapacity() float64 {
	return capacity.NegativeOneSNRCapacity()
}

// HighSNRAssimilation is B*log2(SNR).
func HighSNRAssimilation(b, snrLinear float64) float64 {
	return capacity.HighSNRAssimilation(b, snrLinear)
}

// LimitSaturation verifies the infinite limit does not grow with B.
func LimitSaturation(pOverN0 float64) bool {
	return InfiniteBandwidthLimit(pOverN0) == InfiniteBandwidthLimit(pOverN0)
}

// EfficiencyLowSNR is SNR/ln2.
func EfficiencyLowSNR(snrLinear float64) float64 {
	return snrLinear / math.Ln2
}

// IsLowSNR reports the low-SNR regime.
func IsLowSNR(snrLinear float64) bool {
	return capacity.IsLowSNR(snrLinear)
}

// RelativeLowSNRError returns the approximation relative error.
func RelativeLowSNRError(b, snrLinear float64) float64 {
	exact := capacity.Capacity(b, snrLinear)
	approx := LowSNRCapacity(b, snrLinear)
	if exact == 0 {
		return 0
	}
	return math.Abs(exact-approx) / exact
}

// POverN0FromLimit inverts the infinite-bandwidth limit.
func POverN0FromLimit(cInf float64) float64 {
	return cInf / math.Log2E
}

// SaturationRatio returns C/C_inf.
func SaturationRatio(c, cInf float64) float64 {
	if cInf == 0 {
		return 0
	}
	return c / cInf
}

// BandwidthAtSaturation solves B for a fraction of the infinite limit.
func BandwidthAtSaturation(pOverN0, fraction float64) float64 {
	if fraction <= 0 {
		return 0
	}
	limit := InfiniteBandwidthLimit(pOverN0)
	if limit == 0 {
		return 0
	}
	target := limit * fraction
	// Iteratively widen bandwidth until capacity reaches the target.
	b := 1.0
	for i := 0; i < 80; i++ {
		c := capacity.Capacity(b, pOverN0/b)
		if c >= target {
			return b
		}
		b *= 2
	}
	return b
}
