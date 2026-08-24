package model

import "testing"

func TestValidateRejectsNonPositiveBandwidth(t *testing.T) {
	in := NewCapacityInput(0, 1)
	err := ValidateCapacity(in)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !IsCode(err, CodeInvalidBandwidth) {
		t.Errorf("expected code %q, got %q", CodeInvalidBandwidth, Code(err))
	}
}

func TestValidateRejectsSNRAtNegativeOne(t *testing.T) {
	in := NewCapacityInput(1e6, -1)
	err := ValidateCapacity(in)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !IsCode(err, CodeInvalidSNR) {
		t.Errorf("expected code %q, got %q", CodeInvalidSNR, Code(err))
	}
}

func TestValidateAcceptsDBInput(t *testing.T) {
	in := NewCapacityInput(1e6, 10)
	in.SNRDB = 10
	in.SNRLinear = 10
	if err := ValidateCapacity(in); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateTradeoffBadMode(t *testing.T) {
	in := TradeoffInput{Mode: "bogus"}
	err := ValidateTradeoff(in)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !IsCode(err, CodeInvalidMode) {
		t.Errorf("expected code %q, got %q", CodeInvalidMode, Code(err))
	}
}

func TestMissingField(t *testing.T) {
	if !IsCode(MissingField("b"), CodeMissingField) {
		t.Errorf("missing field code wrong")
	}
}

func TestDBFromLinear(t *testing.T) {
	if got := DBFromLinear(10); got < 9.99 || got > 10.01 {
		t.Errorf("dB = %g, want near 10", got)
	}
}
