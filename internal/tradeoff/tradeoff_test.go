package tradeoff

import (
	"math"
	"testing"

	"shannon-cap/internal/model"
)

func TestRequiredSNRForTargetCapacity(t *testing.T) {
	got := RequiredSNR(10, 10)
	want := math.Exp2(1) - 1
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("SNR = %g, want %g", got, want)
	}
}

func TestRequiredBandwidthForTargetCapacity(t *testing.T) {
	got := RequiredBandwidth(10, 3)
	want := 10 / math.Log2(4)
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("B = %g, want %g", got, want)
	}
}

func TestResolveRequiredSNR(t *testing.T) {
	in := model.TradeoffInput{
		TargetCapacity: 50e6,
		Bandwidth:      10e6,
		Mode:           "required-snr",
	}
	res, err := Resolve(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.RequiredSNRLinear <= 0 {
		t.Errorf("required SNR = %g, want positive", res.RequiredSNRLinear)
	}
}

func TestResolveInfiniteBandwidth(t *testing.T) {
	in := model.TradeoffInput{POverN0: 100, Mode: "infinite-bw"}
	res, err := Resolve(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(res.InfiniteCapacity-100*math.Log2E) > 1e-6 {
		t.Errorf("infinite capacity = %g, want %g", res.InfiniteCapacity, 100*math.Log2E)
	}
}

func TestRequiredPowerSpectralDensity(t *testing.T) {
	got := RequiredPowerSpectralDensity(100 * math.Log2E)
	if math.Abs(got-100) > 1e-9 {
		t.Errorf("P/N0 = %g, want 100", got)
	}
}
