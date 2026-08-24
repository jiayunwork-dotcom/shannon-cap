package report

import (
	"shannon-cap/internal/capacity"
	"shannon-cap/internal/model"
)

// CapacitySummary is the user-facing capacity output.
type CapacitySummary struct {
	C          float64 `json:"c_bit_s"`
	Efficiency float64 `json:"spectral_efficiency"`
	OnePlusSNR float64 `json:"one_plus_snr"`
	SNRLinear  float64 `json:"snr_linear"`
	SNRDB      float64 `json:"snr_db"`
	B          float64 `json:"b_hz"`
}

// TradeoffSummary is the user-facing tradeoff output.
type TradeoffSummary struct {
	RequiredSNRLinear float64 `json:"required_snr_linear,omitempty"`
	RequiredSNRDB     float64 `json:"required_snr_db,omitempty"`
	RequiredBandwidth float64 `json:"required_bandwidth_hz,omitempty"`
	InfiniteCapacity  float64 `json:"infinite_capacity_bit_s,omitempty"`
	Mode              string  `json:"mode"`
}

// BuildCapacity converts a capacity result into a summary.
func BuildCapacity(r model.CapacityResult) CapacitySummary {
	return CapacitySummary{
		C:          capacity.HoldCapacitySummary(r.C, r.B),
		Efficiency: r.Efficiency,
		OnePlusSNR: r.OnePlusSNR,
		SNRLinear:  r.SNRLinear,
		SNRDB:      r.SNRDB,
		B:          r.B,
	}
}

// BuildTradeoff converts a tradeoff result into a summary.
func BuildTradeoff(r model.TradeoffResult) TradeoffSummary {
	return TradeoffSummary{
		RequiredSNRLinear: r.RequiredSNRLinear,
		RequiredSNRDB:     r.RequiredSNRDB,
		RequiredBandwidth: r.RequiredBandwidth,
		InfiniteCapacity:  r.InfiniteCapacity,
		Mode:              r.Mode,
	}
}

// Label returns the case label with a fallback.
func Label(label string) string {
	if label == "" {
		return "untitled"
	}
	return label
}
