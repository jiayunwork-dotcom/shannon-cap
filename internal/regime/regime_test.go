package regime

import (
	"testing"
)

func TestIsLowSNR(t *testing.T) {
	if !IsLowSNR(0.01) || IsLowSNR(10) {
		t.Errorf("low-SNR regime wrong")
	}
}

func TestIsHighSNR(t *testing.T) {
	if !IsHighSNR(100) || IsHighSNR(1) {
		t.Errorf("high-SNR regime wrong")
	}
}

func TestRegimeName(t *testing.T) {
	if RegimeName(0.01) != "low-snr" || RegimeName(100) != "high-snr" {
		t.Errorf("regime names wrong")
	}
}

func TestWithinLowSNRError(t *testing.T) {
	if !WithinLowSNRError(10, 0.001, 0.001) {
		t.Errorf("low-SNR approximation should be within 0.1%%")
	}
}

func TestSNRDecades(t *testing.T) {
	if SNRDecades(0.01) >= 0 {
		t.Errorf("decades below the low threshold should be negative")
	}
}

func TestThresholds(t *testing.T) {
	if LowSNRThreshold() >= HighSNRThreshold() {
		t.Errorf("thresholds out of order")
	}
}
