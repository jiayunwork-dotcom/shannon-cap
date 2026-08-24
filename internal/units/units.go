package units

import (
	"math"
)

const (
	HertzPerMegahertz = 1e6
	HertzPerKilohertz = 1e3
	BitsPerKilobit    = 1e3
	BitsPerMegabit    = 1e6
)

// MHzToHz converts MHz to Hz.
func MHzToHz(mhz float64) float64 {
	return mhz * HertzPerMegahertz
}

// HzToMHz converts Hz to MHz.
func HzToMHz(hz float64) float64 {
	return hz / HertzPerMegahertz
}

// kHzToHz converts kHz to Hz.
func kHzToHz(khz float64) float64 {
	return khz * HertzPerKilohertz
}

// HzTokHz converts Hz to kHz.
func HzTokHz(hz float64) float64 {
	return hz / HertzPerKilohertz
}

// MegabitsToBits converts Mbit/s to bit/s.
func MegabitsToBits(mbit float64) float64 {
	return mbit * BitsPerMegabit
}

// BitsToMegabits converts bit/s to Mbit/s.
func BitsToMegabits(bits float64) float64 {
	return bits / BitsPerMegabit
}

// SNRFromDB converts dB to linear.
func SNRFromDB(db float64) float64 {
	return math.Pow(10, db/10)
}

// DBFromSNR converts linear to dB.
func DBFromSNR(linear float64) float64 {
	if linear <= 0 {
		return -math.MaxFloat64
	}
	return 10 * math.Log10(linear)
}

// IsPositive reports a positive quantity.
func IsPositive(v float64) bool {
	return v > 0
}

// SafeDivide returns zero for a zero denominator.
func SafeDivide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
