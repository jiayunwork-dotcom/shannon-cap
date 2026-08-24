package capacity

import "math"

// LowSNRCapacity returns B*SNR/ln2 for SNR near zero.
func LowSNRCapacity(b, snrLinear float64) float64 {
	return b * snrLinear / math.Ln2
}

// InfiniteBandwidthLimit returns (P/N0)*log2(e).
func InfiniteBandwidthLimit(pOverN0 float64) float64 {
	return pOverN0 * math.Log2E
}

// ZeroSNRCapacity is zero.
func ZeroSNRCapacity() float64 {
	return 0
}

// NegativeOneSNRCapacity is zero for 1+SNR=0.
func NegativeOneSNRCapacity() float64 {
	return 0
}

// HighSNRAssimilation reports C near B*log2(SNR) for large SNR.
func HighSNRAssimilation(b, snrLinear float64) float64 {
	return b * math.Log2(snrLinear)
}

// EfficiencyLowSNR is SNR/ln2.
func EfficiencyLowSNR(snrLinear float64) float64 {
	return snrLinear / math.Ln2
}

// IsLowSNR reports SNR below a threshold.
func IsLowSNR(snrLinear float64) bool {
	return snrLinear < 0.1
}

// PercentOfInfiniteLimit normalizes capacity by the infinite limit.
func PercentOfInfiniteLimit(c, pOverN0 float64) float64 {
	limit := InfiniteBandwidthLimit(pOverN0)
	if limit == 0 {
		return 0
	}
	return c / limit * 100
}
