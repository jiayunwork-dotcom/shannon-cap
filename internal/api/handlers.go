package api

import (
	"net/http"

	"shannon-cap/internal/capacity"
	"shannon-cap/internal/report"
	"shannon-cap/internal/tradeoff"
)

// HandleCapacity computes Shannon capacity for POST /api/capacity.
func HandleCapacity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	in, err := DecodeCapacityBody(LimitBody(r.Body, MaxBodySize))
	if err != nil {
		writeError(w, err)
		return
	}
	res, err := capacity.Compute(in)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report.BuildCapacity(res))
}

// HandleTradeoff computes a bandwidth/power tradeoff for POST /api/tradeoff.
func HandleTradeoff(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	in, err := DecodeTradeoffBody(LimitBody(r.Body, MaxBodySize))
	if err != nil {
		writeError(w, err)
		return
	}
	res, err := tradeoff.Resolve(in)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report.BuildTradeoff(res))
}

// HandleHealth reports service health.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"name":   "shannon-cap",
	})
}

// HandleRoot describes the capacity endpoints.
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"service": "shannon-cap",
		"endpoints": []string{
			"POST /api/capacity",
			"POST /api/tradeoff",
		},
	})
}

// NewRoutes registers every public route.
func NewRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HandleRoot)
	mux.HandleFunc("/health", HandleHealth)
	mux.HandleFunc("/api/capacity", HandleCapacity)
	mux.HandleFunc("/api/tradeoff", HandleTradeoff)
}
