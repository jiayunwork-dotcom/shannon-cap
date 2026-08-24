package bounds

import (
	"testing"

	"shannon-cap/internal/model"
)

func TestCheckCapacityRejectsHugeBandwidth(t *testing.T) {
	in := model.NewCapacityInput(1e18, 1)
	if err := DefaultBounds().CheckCapacity(in); err == nil {
		t.Fatalf("expected bounds error, got nil")
	}
}

func TestCheckTargetCapacity(t *testing.T) {
	if err := DefaultBounds().CheckTargetCapacity(1e20); err == nil {
		t.Fatalf("expected bounds error, got nil")
	}
}

func TestCheckPOverN0(t *testing.T) {
	if err := DefaultBounds().CheckPOverN0(0); err == nil {
		t.Fatalf("expected bounds error, got nil")
	}
}

func TestSummaryNotEmpty(t *testing.T) {
	if DefaultBounds().Summary() == "" {
		t.Errorf("summary should not be empty")
	}
}
