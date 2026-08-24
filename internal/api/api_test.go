package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPICapacitySuccess(t *testing.T) {
	mux := NewMux()
	body := `{"b":10000000,"snr":8,"snr_in_db":true}`
	req := httptest.NewRequest(http.MethodPost, "/api/capacity", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body=%s", rec.Code, rec.Body.String())
	}
	var out map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out["c_bit_s"] == nil || out["spectral_efficiency"] == nil {
		t.Errorf("response missing capacity fields: %v", out)
	}
}

func TestAPITradeoffSuccess(t *testing.T) {
	mux := NewMux()
	body := `{"target_capacity":50000000,"bandwidth":10000000}`
	req := httptest.NewRequest(http.MethodPost, "/api/tradeoff", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body=%s", rec.Code, rec.Body.String())
	}
	var out map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out["required_snr_linear"] == nil {
		t.Errorf("response missing required SNR: %v", out)
	}
}

func TestAPIInvalidInputErrorBody(t *testing.T) {
	mux := NewMux()
	body := `{"b":0,"snr":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/capacity", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400; body=%s", rec.Code, rec.Body.String())
	}
	var out ErrorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out.Error.Code != "invalid_bandwidth" {
		t.Errorf("error code = %q, want invalid_bandwidth", out.Error.Code)
	}
}

func TestAPIInvalidSNRErrorBody(t *testing.T) {
	mux := NewMux()
	body := `{"b":10000000,"snr":-2}`
	req := httptest.NewRequest(http.MethodPost, "/api/capacity", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400; body=%s", rec.Code, rec.Body.String())
	}
	var out ErrorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out.Error.Code != "invalid_snr" {
		t.Errorf("error code = %q, want invalid_snr", out.Error.Code)
	}
}

func TestAPIHealth(t *testing.T) {
	mux := NewMux()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}
