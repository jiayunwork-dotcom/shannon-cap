package capacity

import (
	"math"
)

// CapacityPerHz returns capacity normalized by bandwidth.
func CapacityPerHz(b, snrLinear float64) float64 {
	if b == 0 {
		return 0
	}
	return Capacity(b, snrLinear) / b
}

// CapacityForSNR is an alias for the exact capacity formula.
func CapacityForSNR(b, snrLinear float64) float64 {
	return Capacity(b, snrLinear)
}

// EfficiencyForBandwidth derives spectral efficiency from C and B.
func EfficiencyForBandwidth(c, b float64) float64 {
	if b == 0 {
		return 0
	}
	return c / b
}

// SNRForEfficiency inverts log2(1+SNR).
func SNRForEfficiency(efficiency float64) float64 {
	return math.Exp2(efficiency) - 1
}

// CapacityCurve builds capacity across SNR values.
func CapacityCurve(b float64, snrValues []float64) []float64 {
	out := make([]float64, 0, len(snrValues))
	for _, snr := range snrValues {
		out = append(out, Capacity(b, snr))
	}
	return out
}

// EfficiencyCurve builds spectral efficiency across SNR values.
func EfficiencyCurve(snrValues []float64) []float64 {
	out := make([]float64, 0, len(snrValues))
	for _, snr := range snrValues {
		out = append(out, SpectralEfficiency(snr))
	}
	return out
}

// BandwidthAtCapacity solves B for a target capacity and SNR.
func BandwidthAtCapacity(targetC, snrLinear float64) float64 {
	eta := SpectralEfficiency(snrLinear)
	if eta == 0 {
		return 0
	}
	return targetC / eta
}

// SNRAtCapacity solves SNR for a target capacity and bandwidth.
func SNRAtCapacity(targetC, b float64) float64 {
	if b == 0 {
		return 0
	}
	return math.Exp2(targetC/b) - 1
}

// MaxCapacityAtSNR returns capacity for the supported maximum bandwidth.
func MaxCapacityAtSNR(b, snrLinear float64) float64 {
	return Capacity(b, snrLinear)
}

// CapacityForDB builds capacity from dB.
func CapacityForDB(b, db float64) float64 {
	return Capacity(b, LinearFromDB(db))
}

// EfficiencyForDB builds efficiency from dB.
func EfficiencyForDB(db float64) float64 {
	return SpectralEfficiency(LinearFromDB(db))
}
