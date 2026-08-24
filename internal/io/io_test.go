package io

import (
	"encoding/json"
	"strings"
	"testing"

	"shannon-cap/internal/model"
)

func TestLoadCapacityFileExample(t *testing.T) {
	in, err := LoadCapacityFile("../../example/lte-10mhz.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if in.B != 1e7 || in.SNRDB != 8 {
		t.Errorf("example input wrong: %+v", in)
	}
}

func TestDecodeCapacityRequestMissingField(t *testing.T) {
	body := `{"b":10000000}`
	_, err := DecodeCapacityRequest(strings.NewReader(body))
	if err == nil {
		t.Fatalf("expected missing-field error, got nil")
	}
	if !model.IsCode(err, model.CodeMissingField) {
		t.Errorf("expected code %q, got %q", model.CodeMissingField, model.Code(err))
	}
}

func TestDecodeCapacityRequestDB(t *testing.T) {
	body := `{"b":10000000,"snr":8,"snr_in_db":true}`
	in, err := DecodeCapacityRequest(strings.NewReader(body))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if in.SNRDB != 8 {
		t.Errorf("SNR dB = %g, want 8", in.SNRDB)
	}
}

func TestDecodeTradeoffRequest(t *testing.T) {
	body := `{"target_capacity":50000000,"bandwidth":10000000}`
	in, err := DecodeTradeoffRequest(strings.NewReader(body))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if in.Mode != "required-snr" {
		t.Errorf("mode = %q, want required-snr", in.Mode)
	}
}

func TestWriteJSONRoundTrip(t *testing.T) {
	var buf strings.Builder
	if err := WriteJSON(&buf, map[string]float64{"c": 100}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var out map[string]float64
	if err := json.Unmarshal([]byte(buf.String()), &out); err != nil {
		t.Fatalf("unexpected decode error: %v", err)
	}
	if out["c"] != 100 {
		t.Errorf("decoded c = %g, want 100", out["c"])
	}
}

func TestDescribeInput(t *testing.T) {
	in := model.NewCapacityInput(10e6, 8)
	if DescribeInput(in) == "" {
		t.Errorf("description should not be empty")
	}
}
