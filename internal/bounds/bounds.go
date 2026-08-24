package bounds

import (
	"math"

	"shannon-cap/internal/model"
)

// Bounds keeps channel quantities inside a defensive operating envelope.
type Bounds struct {
	MaxBandwidth float64
	MaxSNRLinear float64
	MaxCapacity  float64
	MaxPOverN0   float64
}

// DefaultBounds returns the standard defensive limits.
func DefaultBounds() Bounds {
	return Bounds{
		MaxBandwidth: 1e15,
		MaxSNRLinear: 1e18,
		MaxCapacity:  1e18,
		MaxPOverN0:   1e18,
	}
}

// CheckCapacity returns the first bound violation, if any.
func (b Bounds) CheckCapacity(in model.CapacityInput) error {
	if in.B > b.MaxBandwidth {
		return model.NewError(model.CodeOutOfRange, "bandwidth exceeds the supported bound").
			WithField("b", in.B)
	}
	if in.SNRLinear > b.MaxSNRLinear {
		return model.NewError(model.CodeOutOfRange, "linear SNR exceeds the supported bound").
			WithField("snr", in.SNRLinear)
	}
	return nil
}

// CheckTargetCapacity validates the target against the envelope.
func (b Bounds) CheckTargetCapacity(c float64) error {
	if c > b.MaxCapacity {
		return model.NewError(model.CodeInvalidTarget, "target capacity exceeds the supported bound").
			WithField("target_capacity", c)
	}
	return nil
}

// CheckPOverN0 validates the power spectral density against the envelope.
func (b Bounds) CheckPOverN0(p float64) error {
	if math.IsNaN(p) || math.IsInf(p, 0) || p <= 0 || p > b.MaxPOverN0 {
		return model.NewError(model.CodeInvalidMode, "P/N0 must be positive and within the supported bound").
			WithField("p_over_n0", p)
	}
	return nil
}

// Summary returns a human-readable envelope.
func (b Bounds) Summary() string {
	return "B<=1e15 SNR<=1e18 C<=1e18 P/N0<=1e18"
}
