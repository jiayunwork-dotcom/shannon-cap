package io

import (
	"encoding/json"
	"io"
	"os"

	"shannon-cap/internal/capacity"
	"shannon-cap/internal/model"
)

// CapacityRequest is the JSON shape accepted by the capacity endpoint.
type CapacityRequest struct {
	B       *float64 `json:"b"`
	SNR     *float64 `json:"snr"`
	SNRInDB *bool    `json:"snr_in_db,omitempty"`
	Label   string   `json:"label,omitempty"`
}

// TradeoffRequest is the JSON shape accepted by the tradeoff endpoint.
type TradeoffRequest struct {
	TargetCapacity *float64 `json:"target_capacity,omitempty"`
	Bandwidth      *float64 `json:"bandwidth,omitempty"`
	SNR            *float64 `json:"snr,omitempty"`
	SNRInDB        *bool    `json:"snr_in_db,omitempty"`
	POverN0        *float64 `json:"p_over_n0,omitempty"`
	Label          string   `json:"label,omitempty"`
}

// ToInput converts a validated capacity request into a model input.
func (r CapacityRequest) ToInput() (model.CapacityInput, error) {
	if r.B == nil {
		return model.CapacityInput{}, model.MissingField("b")
	}
	if r.SNR == nil {
		return model.CapacityInput{}, model.MissingField("snr")
	}
	inDB := r.SNRInDB != nil && *r.SNRInDB
	if inDB {
		return model.CapacityInput{
			B:         *r.B,
			SNRLinear: capacity.LinearFromDB(*r.SNR),
			SNRDB:     *r.SNR,
			Label:     r.Label,
		}, nil
	}
	return model.CapacityInput{
		B:         *r.B,
		SNRLinear: *r.SNR,
		SNRDB:     capacity.DBFromLinear(*r.SNR),
		Label:     r.Label,
	}, nil
}

// ToInput converts a validated tradeoff request into a model input.
func (r TradeoffRequest) ToInput() (model.TradeoffInput, error) {
	if r.POverN0 != nil {
		return model.TradeoffInput{
			POverN0: *r.POverN0,
			Mode:    "infinite-bw",
			Label:   r.Label,
		}, nil
	}
	if r.TargetCapacity == nil {
		return model.TradeoffInput{}, model.MissingField("target_capacity")
	}
	if r.Bandwidth != nil {
		return model.TradeoffInput{
			TargetCapacity: *r.TargetCapacity,
			Bandwidth:      *r.Bandwidth,
			Mode:           "required-snr",
			Label:          r.Label,
		}, nil
	}
	if r.SNR == nil {
		return model.TradeoffInput{}, model.MissingField("snr")
	}
	inDB := r.SNRInDB != nil && *r.SNRInDB
	if inDB {
		return model.TradeoffInput{
			TargetCapacity: *r.TargetCapacity,
			SNRLinear:      capacity.LinearFromDB(*r.SNR),
			SNRDB:          *r.SNR,
			Mode:           "required-bw",
			Label:          r.Label,
		}, nil
	}
	return model.TradeoffInput{
		TargetCapacity: *r.TargetCapacity,
		SNRLinear:      *r.SNR,
		SNRDB:          capacity.DBFromLinear(*r.SNR),
		Mode:           "required-bw",
		Label:          r.Label,
	}, nil
}

// DecodeCapacityRequest reads and converts a capacity JSON body.
func DecodeCapacityRequest(r io.Reader) (model.CapacityInput, error) {
	var req CapacityRequest
	if err := json.NewDecoder(r).Decode(&req); err != nil {
		return model.CapacityInput{}, model.NewError(model.CodeInvalidJSON, err.Error())
	}
	return req.ToInput()
}

// DecodeTradeoffRequest reads and converts a tradeoff JSON body.
func DecodeTradeoffRequest(r io.Reader) (model.TradeoffInput, error) {
	var req TradeoffRequest
	if err := json.NewDecoder(r).Decode(&req); err != nil {
		return model.TradeoffInput{}, model.NewError(model.CodeInvalidJSON, err.Error())
	}
	return req.ToInput()
}

// LoadCapacityFile reads a capacity request from disk.
func LoadCapacityFile(path string) (model.CapacityInput, error) {
	f, err := os.Open(path)
	if err != nil {
		return model.CapacityInput{}, err
	}
	defer f.Close()
	return DecodeCapacityRequest(f)
}

// LoadTradeoffFile reads a tradeoff request from disk.
func LoadTradeoffFile(path string) (model.TradeoffInput, error) {
	f, err := os.Open(path)
	if err != nil {
		return model.TradeoffInput{}, err
	}
	defer f.Close()
	return DecodeTradeoffRequest(f)
}

// ExamplePath returns the bundled LTE example.
func ExamplePath() string {
	return "example/lte-10mhz.json"
}
