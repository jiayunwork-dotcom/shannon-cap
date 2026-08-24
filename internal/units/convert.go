package units

import "shannon-cap/internal/model"

// CapacityInputFromMHzAndDB builds an input from MHz and dB.
func CapacityInputFromMHzAndDB(mhz, db float64) model.CapacityInput {
	return model.CapacityInput{
		B:         MHzToHz(mhz),
		SNRLinear: SNRFromDB(db),
		SNRDB:     db,
	}
}

// CapacityInputFromKHz builds an input from kHz and linear SNR.
func CapacityInputFromKHz(khz, snrLinear float64) model.CapacityInput {
	return model.CapacityInput{
		B:         kHzToHz(khz),
		SNRLinear: snrLinear,
		SNRDB:     DBFromSNR(snrLinear),
	}
}

// ConvertCapacityToMbits converts bit/s to Mbit/s.
func ConvertCapacityToMbits(bits float64) float64 {
	return BitsToMegabits(bits)
}

// ConvertBandwidthToMHz converts Hz to MHz.
func ConvertBandwidthToMHz(hz float64) float64 {
	return HzToMHz(hz)
}

// UnitLabel returns a display label for a bandwidth unit.
func UnitLabel(unit string) string {
	switch unit {
	case "MHz":
		return "MHz"
	case "kHz":
		return "kHz"
	default:
		return "Hz"
	}
}

// NormalizeUnit accepts common bandwidth unit aliases.
func NormalizeUnit(unit string) string {
	switch unit {
	case "MHZ", "mhz":
		return "MHz"
	case "KHZ", "khz":
		return "kHz"
	default:
		return "Hz"
	}
}

// HasValidUnit reports whether a unit string is recognized.
func HasValidUnit(unit string) bool {
	return NormalizeUnit(unit) != ""
}

// ToHz converts a quantity to Hz using the unit label.
func ToHz(value float64, unit string) float64 {
	switch NormalizeUnit(unit) {
	case "MHz":
		return MHzToHz(value)
	case "kHz":
		return kHzToHz(value)
	default:
		return value
	}
}

// FromHz converts Hz to the requested unit.
func FromHz(value float64, unit string) float64 {
	switch NormalizeUnit(unit) {
	case "MHz":
		return HzToMHz(value)
	case "kHz":
		return HzTokHz(value)
	default:
		return value
	}
}
