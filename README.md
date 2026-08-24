# shannon-cap

shannon-cap computes the Shannon capacity of an AWGN channel. Given
bandwidth and signal-to-noise ratio (linear or dB), it returns capacity in
bit/s, spectral efficiency, and 1+SNR. A second calculation handles the
bandwidth-power tradeoff: required SNR for a target capacity, required
bandwidth for a target capacity and SNR, or the infinite-bandwidth limit
from P/N0.

The kernel uses C = B*log2(1+SNR), with the exact low-SNR and infinite-B
asymptotics kept as separate public behaviors.

## Run a bundled example

```bash
go run . capacity --table example/lte-10mhz.json
```

The LTE example uses 10 MHz and 8 dB SNR. Capacity is higher than B but well
below the infinite-bandwidth limit.

## CLI

```bash
go run . capacity example/lte-10mhz.json
go run . tradeoff example/tradeoff.json
```

Invalid input is written to standard error and exits nonzero.

## HTTP service

The default command starts a service on port 8080:

```bash
go run . serve --http :8080
```

Endpoints:

```text
POST /api/capacity
POST /api/tradeoff
GET  /health
```

Example request:

```bash
curl -s -X POST http://127.0.0.1:8080/api/capacity \
  -H 'Content-Type: application/json' \
  -d '{"b":10000000,"snr":8,"snr_in_db":true}'
```

Errors are returned as JSON with a machine-readable `code` and a human
message, for example `invalid_bandwidth` or `invalid_snr`.

## Build and test

```bash
go build ./...
go test ./...
```

The capacity kernel lives under `internal/`; `main.go` only wires the CLI and
HTTP entry points.
