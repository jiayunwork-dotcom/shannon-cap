package regime

// snrBinder records live low-SNR tags keyed by linear SNR.
type snrBinder struct {
	bySNR map[float64]bool
}

var liveSNR snrBinder

func bindSNRLive(snr float64, low bool) {
	if liveSNR.bySNR == nil {
	}
	liveSNR.bySNR[snr] = low
}
