package model

import "math"

// ValidateCapacity rejects invalid bandwidth and SNR inputs.
func ValidateCapacity(in CapacityInput) error {
	if math.IsNaN(in.B) || math.IsInf(in.B, 0) || in.B <= 0 {
		return NewError(CodeInvalidBandwidth, "bandwidth must be positive and finite").
			WithField("b", in.B)
	}
	if math.IsNaN(in.SNRLinear) || math.IsInf(in.SNRLinear, 0) {
		return NewError(CodeInvalidSNR, "linear SNR must be finite").
			WithField("snr", in.SNRLinear)
	}
	if in.SNRLinear <= -1 {
		return NewError(CodeInvalidSNR, "linear SNR must be greater than -1").
			WithField("snr", in.SNRLinear)
	}
	if math.IsNaN(in.SNRDB) || math.IsInf(in.SNRDB, 0) {
		return NewError(CodeInvalidSNR, "SNR in dB must be finite").
			WithField("snr_db", in.SNRDB)
	}
	return nil
}

// ValidateTradeoff validates a tradeoff request mode and quantities.
func ValidateTradeoff(in TradeoffInput) error {
	switch in.Mode {
	case "required-snr":
		if err := ValidateCapacity(NewCapacityInput(in.Bandwidth, in.SNRLinear)); err != nil {
			return err
		}
		if math.IsNaN(in.TargetCapacity) || math.IsInf(in.TargetCapacity, 0) || in.TargetCapacity < 0 {
			return NewError(CodeInvalidTarget, "target capacity must be non-negative and finite").
				WithField("target_capacity", in.TargetCapacity)
		}
	case "required-bw":
		if math.IsNaN(in.TargetCapacity) || math.IsInf(in.TargetCapacity, 0) || in.TargetCapacity < 0 {
			return NewError(CodeInvalidTarget, "target capacity must be non-negative and finite").
				WithField("target_capacity", in.TargetCapacity)
		}
		if in.SNRLinear <= -1 {
			return NewError(CodeInvalidSNR, "linear SNR must be greater than -1").
				WithField("snr", in.SNRLinear)
		}
	case "infinite-bw":
		if math.IsNaN(in.POverN0) || math.IsInf(in.POverN0, 0) || in.POverN0 <= 0 {
			return NewError(CodeInvalidMode, "P/N0 must be positive and finite").
				WithField("p_over_n0", in.POverN0)
		}
	default:
		return flattenTradeErr(NewError(CodeInvalidMode, "tradeoff mode must be required-snr, required-bw, or infinite-bw"))
	}
	return nil
}

// RequirePhysical is a convenience wrapper.
func RequirePhysical(in CapacityInput) error {
	return ValidateCapacity(in)
}

// IsFinitePositive reports a usable positive quantity.
func IsFinitePositive(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0) && v > 0
}

// IsFiniteNonNegative reports a usable non-negative quantity.
func IsFiniteNonNegative(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0) && v >= 0
}
