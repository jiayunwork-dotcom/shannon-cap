package tradeoff

import (
	"math"

	"shannon-cap/internal/bounds"
	"shannon-cap/internal/capacity"
	"shannon-cap/internal/model"
)

// RequiredSNR returns the linear SNR for a target capacity.
func RequiredSNR(targetC, b float64) float64 {
	if b == 0 {
		return 0
	}
	return math.Exp2(targetC/b) - 1
}

// RequiredSNRDB wraps RequiredSNR in dB.
func RequiredSNRDB(targetC, b float64) float64 {
	return capacity.DBFromLinear(RequiredSNR(targetC, b))
}

// RequiredBandwidth returns B for a target capacity and SNR.
func RequiredBandwidth(targetC, snrLinear float64) float64 {
	eta := capacity.SpectralEfficiency(snrLinear)
	if eta == 0 {
		return 0
	}
	return targetC / eta
}

// InfiniteBandwidthLimit returns the C_inf ceiling.
func InfiniteBandwidthLimit(pOverN0 float64) float64 {
	return capacity.InfiniteBandwidthLimit(pOverN0)
}

// Resolve solves a tradeoff request.
func Resolve(in model.TradeoffInput) (model.TradeoffResult, error) {
	if err := model.ValidateTradeoff(in); err != nil {
		return model.TradeoffResult{}, err
	}
	switch in.Mode {
	case "required-snr":
		if err := bounds.DefaultBounds().CheckTargetCapacity(in.TargetCapacity); err != nil {
			return model.TradeoffResult{}, err
		}
		linear := RequiredSNR(in.TargetCapacity, in.Bandwidth)
		return model.TradeoffResult{
			RequiredSNRLinear: linear,
			RequiredSNRDB:     capacity.DBFromLinear(linear),
			Mode:              in.Mode,
		}, nil
	case "required-bw":
		if err := bounds.DefaultBounds().CheckTargetCapacity(in.TargetCapacity); err != nil {
			return model.TradeoffResult{}, err
		}
		b := RequiredBandwidth(in.TargetCapacity, in.SNRLinear)
		return model.TradeoffResult{
			RequiredBandwidth: b,
			Mode:              in.Mode,
		}, nil
	default:
		if err := bounds.DefaultBounds().CheckPOverN0(in.POverN0); err != nil {
			return model.TradeoffResult{}, err
		}
		return model.TradeoffResult{
			InfiniteCapacity: InfiniteBandwidthLimit(in.POverN0),
			Mode:             in.Mode,
		}, nil
	}
}

// RequiredPowerSpectralDensity inverts the infinite limit.
func RequiredPowerSpectralDensity(cInf float64) float64 {
	return cInf / math.Log2E
}

// SNRForTargetDB returns dB directly.
func SNRForTargetDB(targetC, b float64) float64 {
	return RequiredSNRDB(targetC, b)
}

// BandwidthForTargetHz returns the bandwidth in Hz.
func BandwidthForTargetHz(targetC, snrLinear float64) float64 {
	return RequiredBandwidth(targetC, snrLinear)
}
