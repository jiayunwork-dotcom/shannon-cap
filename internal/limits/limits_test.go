package limits

import (
	"math"
	"testing"
)

func TestLowSNRCapacity(t *testing.T) {
	got := LowSNRCapacity(10, 0.01)
	want := 10 * 0.01 / math.Ln2
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("capacity = %g, want %g", got, want)
	}
}

func TestInfiniteBandwidthLimit(t *testing.T) {
	got := InfiniteBandwidthLimit(100)
	if math.Abs(got-100*math.Log2E) > 1e-9 {
		t.Errorf("limit = %g, want %g", got, 100*math.Log2E)
	}
}

func TestZeroSNRCapacity(t *testing.T) {
	if ZeroSNRCapacity() != 0 {
		t.Errorf("zero-SNR capacity should be zero")
	}
}

func TestNegativeOneSNRCapacity(t *testing.T) {
	if NegativeOneSNRCapacity() != 0 {
		t.Errorf("SNR=-1 capacity should be zero")
	}
}

func TestPOverN0FromLimit(t *testing.T) {
	got := POverN0FromLimit(100 * math.Log2E)
	if math.Abs(got-100) > 1e-9 {
		t.Errorf("P/N0 = %g, want 100", got)
	}
}

func TestLimitSaturation(t *testing.T) {
	if !LimitSaturation(100) {
		t.Errorf("infinite limit should be independent of B")
	}
}
