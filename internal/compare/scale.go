package compare

// ScaleCase pairs a bandwidth multiplier with an expected capacity multiplier.
type ScaleCase struct {
	InputMultiplier  float64
	OutputMultiplier float64
}

// BandwidthCases returns the bandwidth scaling cases.
func BandwidthCases() []ScaleCase {
	return []ScaleCase{
		{InputMultiplier: 1, OutputMultiplier: 1},
		{InputMultiplier: 2, OutputMultiplier: 2},
		{InputMultiplier: 0.5, OutputMultiplier: 0.5},
	}
}

// MatchScale verifies a measured multiplier.
func MatchScale(c ScaleCase, actual float64) bool {
	return WithinRelative(actual, c.OutputMultiplier, 1e-9)
}

// RelativeError returns the unsigned relative difference.
func RelativeError(actual, expected float64) float64 {
	if expected == 0 {
		if actual == 0 {
			return 0
		}
		return 1
	}
	diff := actual - expected
	if diff < 0 {
		diff = -diff
	}
	if expected < 0 {
		expected = -expected
	}
	return diff / expected
}

// WithinRelative reports whether a relative error is acceptable.
func WithinRelative(actual, expected, maxErr float64) bool {
	return RelativeError(actual, expected) <= maxErr
}

// PercentDifference formats a relative error.
func PercentDifference(actual, expected float64) float64 {
	return RelativeError(actual, expected) * 100
}
