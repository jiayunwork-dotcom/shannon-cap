package compare

import "testing"

func TestZeroSNRInvariant(t *testing.T) {
	if !ZeroSNRInvariant(10) {
		t.Errorf("SNR=0 should give zero capacity")
	}
}

func TestBandwidthDoublingInvariant(t *testing.T) {
	if !BandwidthDoublingInvariant(10, 3) {
		t.Errorf("doubling B should double C")
	}
}

func TestDBIncreaseUnderTenfold(t *testing.T) {
	if !DBIncreaseUnderTenfold(10) {
		t.Errorf("0 dB to 10 dB should raise C but not 10x")
	}
}

func TestInfiniteLimitSaturation(t *testing.T) {
	if !InfiniteLimitSaturation(100) {
		t.Errorf("infinite limit should not grow with B")
	}
}

func TestNegativeOneSNRInvariant(t *testing.T) {
	if !NegativeOneSNRInvariant() {
		t.Errorf("SNR=-1 should give zero capacity")
	}
}

func TestLowSNRApproximation(t *testing.T) {
	if !LowSNRApproximation(10, 0.001) {
		t.Errorf("low-SNR approximation should be close")
	}
}

func TestSameSNRConversion(t *testing.T) {
	if !SameSNRConversion(12.5) {
		t.Errorf("linear/dB conversion should round trip")
	}
}

func TestMatchScale(t *testing.T) {
	c := ScaleCase{InputMultiplier: 2, OutputMultiplier: 2}
	if !MatchScale(c, 2) {
		t.Errorf("bandwidth scaling case should match")
	}
}
