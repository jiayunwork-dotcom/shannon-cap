package tradeoff

import (
	"math"

	"shannon-cap/internal/capacity"
)

// CurvePoint pairs a bandwidth with its capacity.
type CurvePoint struct {
	B float64
	C float64
}

// TradeoffCurve builds capacity as a function of bandwidth.
func TradeoffCurve(pOverN0 float64, bandwidths []float64) []CurvePoint {
	out := make([]CurvePoint, 0, len(bandwidths))
	for _, b := range bandwidths {
		out = append(out, CurvePoint{
			B: b,
			C: capacity.Capacity(b, pOverN0/b),
		})
	}
	return out
}

// SNRCurve returns the SNR implied by each bandwidth.
func SNRCurve(pOverN0 float64, bandwidths []float64) []float64 {
	out := make([]float64, 0, len(bandwidths))
	for _, b := range bandwidths {
		out = append(out, pOverN0/b)
	}
	return out
}

// BandwidthCurve returns bandwidths spanning decades.
func BandwidthCurve(start, end float64, count int) []float64 {
	if count < 2 {
		count = 2
	}
	out := make([]float64, 0, count)
	logStart := math.Log10(start)
	logEnd := math.Log10(end)
	for i := 0; i < count; i++ {
		logB := logStart + (logEnd-logStart)*float64(i)/float64(count-1)
		out = append(out, math.Pow(10, logB))
	}
	return out
}

// CurveMaxCapacity returns the largest capacity in a curve.
func CurveMaxCapacity(points []CurvePoint) float64 {
	max := 0.0
	for _, p := range points {
		if p.C > max {
			max = p.C
		}
	}
	return max
}

// CurveSaturation reports the fraction of the infinite limit at the last point.
func CurveSaturation(points []CurvePoint, pOverN0 float64) float64 {
	if len(points) == 0 {
		return 0
	}
	last := points[len(points)-1]
	return capacity.PercentOfInfiniteLimit(last.C, pOverN0) / 100
}

// RequiredSNRSeries returns SNR values for a target capacity series.
func RequiredSNRSeries(targets []float64, b float64) []float64 {
	out := make([]float64, 0, len(targets))
	for _, c := range targets {
		out = append(out, RequiredSNR(c, b))
	}
	return out
}

// RequiredBandwidthSeries returns bandwidth values for a target series.
func RequiredBandwidthSeries(targets []float64, snrLinear float64) []float64 {
	out := make([]float64, 0, len(targets))
	for _, c := range targets {
		out = append(out, RequiredBandwidth(c, snrLinear))
	}
	return out
}

// CurveLength returns the number of points.
func CurveLength(points []CurvePoint) int {
	return len(points)
}

// PercentOfLimitAt returns the limit percentage at a point.
func PercentOfLimitAt(p CurvePoint, pOverN0 float64) float64 {
	return capacity.PercentOfInfiniteLimit(p.C, pOverN0)
}
