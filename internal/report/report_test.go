package report

import (
	"testing"

	"shannon-cap/internal/model"
)

func TestBuildCapacitySummary(t *testing.T) {
	res := model.CapacityResult{C: 100, Efficiency: 2, OnePlusSNR: 4, SNRLinear: 3, SNRDB: 6, B: 50}
	s := BuildCapacity(res)
	if s.C != 100 || s.Efficiency != 2 || s.B != 50 {
		t.Errorf("summary fields wrong: %+v", s)
	}
}

func TestBuildTradeoffSummary(t *testing.T) {
	res := model.TradeoffResult{RequiredSNRLinear: 5, Mode: "required-snr"}
	s := BuildTradeoff(res)
	if s.RequiredSNRLinear != 5 || s.Mode != "required-snr" {
		t.Errorf("summary fields wrong: %+v", s)
	}
}

func TestFormatNumber(t *testing.T) {
	if got := FormatNumber(10e6); got != "1.000000e+07" {
		t.Errorf("format = %q, want scientific", got)
	}
}

func TestLabelFallback(t *testing.T) {
	if Label("") != "untitled" {
		t.Errorf("label fallback wrong")
	}
}
