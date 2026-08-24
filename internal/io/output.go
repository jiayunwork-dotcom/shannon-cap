package io

import (
	"encoding/json"
	"io"

	"shannon-cap/internal/model"
	"shannon-cap/internal/report"
)

// WriteJSON encodes a value with indentation.
func WriteJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// WriteCapacityOutput writes a capacity summary.
func WriteCapacityOutput(w io.Writer, s report.CapacitySummary) error {
	return WriteJSON(w, s)
}

// WriteTradeoffOutput writes a tradeoff summary.
func WriteTradeoffOutput(w io.Writer, s report.TradeoffSummary) error {
	return WriteJSON(w, s)
}

// WriteModelError writes a structured error body.
func WriteModelError(w io.Writer, err error) error {
	body := map[string]any{
		"error": map[string]any{
			"code":    model.Code(err),
			"message": err.Error(),
		},
	}
	return WriteJSON(w, body)
}

// WriteCapacityTable writes a human-readable capacity table.
func WriteCapacityTable(w io.Writer, s report.CapacitySummary) error {
	t := &report.Table{}
	t.AddFloat("capacity (bit/s)", s.C)
	t.AddFloat("spectral efficiency", s.Efficiency)
	t.AddFloat("1+SNR", s.OnePlusSNR)
	t.AddFloat("SNR linear", s.SNRLinear)
	t.AddFloat("SNR dB", s.SNRDB)
	t.AddFloat("bandwidth (Hz)", s.B)
	return t.Render(w)
}

// WriteTradeoffTable writes a human-readable tradeoff table.
func WriteTradeoffTable(w io.Writer, s report.TradeoffSummary) error {
	t := &report.Table{}
	t.AddFloat("required SNR linear", s.RequiredSNRLinear)
	t.AddFloat("required SNR dB", s.RequiredSNRDB)
	t.AddFloat("required bandwidth (Hz)", s.RequiredBandwidth)
	t.AddFloat("infinite capacity (bit/s)", s.InfiniteCapacity)
	t.Add("mode", s.Mode)
	return t.Render(w)
}
