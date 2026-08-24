package units

import (
	"math"
	"testing"
)

func TestMHzToHz(t *testing.T) {
	if math.Abs(MHzToHz(10)-1e7) > 1e-9 {
		t.Errorf("MHz conversion wrong")
	}
}

func TestHzToMHz(t *testing.T) {
	if math.Abs(HzToMHz(1e7)-10) > 1e-9 {
		t.Errorf("Hz conversion wrong")
	}
}

func TestMegabitsToBits(t *testing.T) {
	if math.Abs(MegabitsToBits(5)-5e6) > 1e-9 {
		t.Errorf("Mbit conversion wrong")
	}
}

func TestSNRFromDB(t *testing.T) {
	if math.Abs(SNRFromDB(10)-10) > 1e-9 {
		t.Errorf("dB conversion wrong")
	}
}

func TestDBFromSNR(t *testing.T) {
	if math.Abs(DBFromSNR(100)-20) > 1e-9 {
		t.Errorf("linear conversion wrong")
	}
}

func TestSafeDivide(t *testing.T) {
	if SafeDivide(1, 0) != 0 {
		t.Errorf("safe divide wrong")
	}
}
