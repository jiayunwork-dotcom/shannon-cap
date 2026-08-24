package report

import (
	"fmt"
	"math"
	"strings"
)

// FormatNumber prints a float with enough digits.
func FormatNumber(v float64) string {
	if math.Abs(v) >= 1e5 || (math.Abs(v) < 1e-4 && v != 0) {
		return fmt.Sprintf("%.6e", v)
	}
	return fmt.Sprintf("%.8g", v)
}

// FormatCompact trims trailing zeros.
func FormatCompact(v float64) string {
	s := fmt.Sprintf("%.6g", v)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

// FormatCapacity prints a capacity in bit/s.
func FormatCapacity(v float64) string {
	return FormatNumber(v) + " bit/s"
}

// FormatSNR prints a dB SNR.
func FormatSNR(v float64) string {
	return FormatNumber(v) + " dB"
}

// FormatBandwidth prints a bandwidth in Hz.
func FormatBandwidth(v float64) string {
	return FormatNumber(v) + " Hz"
}

// PadRight right-pads a label.
func PadRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
