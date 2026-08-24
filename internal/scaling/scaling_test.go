package scaling

import (
	"math"
	"testing"
)

func TestDoubleBandwidth(t *testing.T) {
	got := DoubleBandwidth(10, 3)
	want := 20 * math.Log2(4)
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("capacity = %g, want %g", got, want)
	}
}

func TestCapacityRatio(t *testing.T) {
	if math.Abs(CapacityRatio(20, 10)-2) > 1e-12 {
		t.Errorf("ratio wrong")
	}
}

func TestPowerScaling(t *testing.T) {
	got := PowerScaling(100, 2)
	want := 200 * math.Log2E
	if math.Abs(got-want) > 1e-6 {
		t.Errorf("power scaling wrong")
	}
}

func TestSNRScale(t *testing.T) {
	if math.Abs(SNRScale(10)-10) > 1e-9 {
		t.Errorf("dB step wrong")
	}
}

func TestCapacityAfterSNRStep(t *testing.T) {
	got := CapacityAfterSNRStep(10, 1, 10)
	want := 10 * math.Log2(11)
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("capacity after step wrong")
	}
}

func TestTenDecibelsFactor(t *testing.T) {
	if TenDecibelsFactor() != 10 {
		t.Errorf("10 dB factor wrong")
	}
}
