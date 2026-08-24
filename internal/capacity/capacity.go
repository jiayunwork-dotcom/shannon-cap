package capacity

import (
	"math"

	"shannon-cap/internal/bounds"
	"shannon-cap/internal/model"
)

// Capacity computes C = B*log2(1+SNR).
func Capacity(b, snrLinear float64) float64 {
	if snrLinear <= -1 {
		return 0
	}
	return b * math.Log2(1+snrLinear)
}

// SpectralEfficiency returns C/B = log2(1+SNR).
func SpectralEfficiency(snrLinear float64) float64 {
	if snrLinear <= -1 {
		return 0
	}
	return math.Log2(1 + snrLinear)
}

// OnePlusSNR returns 1+SNR.
func OnePlusSNR(snrLinear float64) float64 {
	return 1 + snrLinear
}

// LinearFromDB converts a dB SNR to linear.
func LinearFromDB(db float64) float64 {
	return math.Pow(10, db/10)
}

// DBFromLinear converts a linear SNR to dB.
func DBFromLinear(linear float64) float64 {
	if linear <= 0 {
		return -math.MaxFloat64
	}
	return 10 * math.Log10(linear)
}

// Compute validates and solves one capacity case.
func Compute(in model.CapacityInput) (model.CapacityResult, error) {
	if err := model.ValidateCapacity(in); err != nil {
		return model.CapacityResult{}, err
	}
	if err := boundsCheck(in); err != nil {
		return model.CapacityResult{}, err
	}
	c := Capacity(in.B, in.SNRLinear)
	return model.CapacityResult{
		C:          c,
		Efficiency: SpectralEfficiency(in.SNRLinear),
		OnePlusSNR: OnePlusSNR(in.SNRLinear),
		SNRLinear:  in.SNRLinear,
		SNRDB:      in.SNRDB,
		B:          in.B,
	}, nil
}

// CapacityFromDB computes capacity from a dB SNR.
func CapacityFromDB(b, db float64) float64 {
	return Capacity(b, LinearFromDB(db))
}

// EfficiencyFromDB returns spectral efficiency from dB.
func EfficiencyFromDB(db float64) float64 {
	return SpectralEfficiency(LinearFromDB(db))
}

// BitsPerSecondLabel returns a stable unit label.
func BitsPerSecondLabel() string {
	return "bit/s"
}

func boundsCheck(in model.CapacityInput) error {
	return bounds.DefaultBounds().CheckCapacity(in)
}
