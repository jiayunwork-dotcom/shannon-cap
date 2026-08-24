package io

import (
	"io"

	"shannon-cap/internal/capacity"
	"shannon-cap/internal/model"
	"shannon-cap/internal/report"
	"shannon-cap/internal/tradeoff"
)

// RunCapacity computes a capacity case from a JSON file.
func RunCapacity(path string, table bool, out io.Writer) error {
	in, err := LoadCapacityFile(path)
	if err != nil {
		return err
	}
	res, err := capacity.Compute(in)
	if err != nil {
		return err
	}
	summary := report.BuildCapacity(res)
	if table {
		return WriteCapacityTable(out, summary)
	}
	return WriteCapacityOutput(out, summary)
}

// RunTradeoff computes a tradeoff case from a JSON file.
func RunTradeoff(path string, table bool, out io.Writer) error {
	in, err := LoadTradeoffFile(path)
	if err != nil {
		return err
	}
	res, err := tradeoff.Resolve(in)
	if err != nil {
		return err
	}
	summary := report.BuildTradeoff(res)
	if table {
		return WriteTradeoffTable(out, summary)
	}
	return WriteTradeoffOutput(out, summary)
}

// BuildCapacityFromValues constructs a capacity input from CLI values.
func BuildCapacityFromValues(b, snr float64, inDB bool) (model.CapacityInput, error) {
	if inDB {
		in := model.CapacityInput{B: b, SNRDB: snr, SNRLinear: capacity.LinearFromDB(snr)}
		if err := model.ValidateCapacity(in); err != nil {
			return model.CapacityInput{}, err
		}
		return in, nil
	}
	return model.NewCapacityInput(b, snr), nil
}

// DescribeInput formats the capacity inputs.
func DescribeInput(in model.CapacityInput) string {
	return "B=" + report.FormatNumber(in.B) + " Hz SNR=" + report.FormatNumber(in.SNRLinear) +
		" (" + report.FormatNumber(in.SNRDB) + " dB)"
}
