package capacity

import (
	"math"
	"testing"

	"shannon-cap/internal/model"
)

func TestCapacityExactFormula(t *testing.T) {
	got := Capacity(10, 3)
	want := 10 * math.Log2(4)
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("C = %g, want %g", got, want)
	}
}

func TestZeroSNRCapacity(t *testing.T) {
	if got := Capacity(10, 0); got != 0 {
		t.Errorf("C = %g, want 0", got)
	}
}

func TestNegativeOneSNRCapacity(t *testing.T) {
	if got := Capacity(10, -1); got != 0 {
		t.Errorf("C = %g, want 0", got)
	}
}

func TestLowSNRApproximation(t *testing.T) {
	exact := Capacity(10, 0.001)
	approx := LowSNRCapacity(10, 0.001)
	if math.Abs(exact-approx) > 1e-3*exact {
		t.Errorf("low-SNR approximation too far: exact=%g approx=%g", exact, approx)
	}
}

func TestInfiniteBandwidthLimit(t *testing.T) {
	got := InfiniteBandwidthLimit(100)
	want := 100 * math.Log2E
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("limit = %g, want %g", got, want)
	}
}

func TestLinearDBRoundTrip(t *testing.T) {
	linear := 12.5
	if math.Abs(LinearFromDB(DBFromLinear(linear))-linear) > 1e-12*linear {
		t.Errorf("linear/dB round trip failed")
	}
}

func TestComputeCapacity(t *testing.T) {
	in := model.NewCapacityInput(10, 3)
	res, err := Compute(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.C <= 0 || res.Efficiency <= 0 {
		t.Errorf("result should be positive: %+v", res)
	}
}
