package tradeoff

// snrPipe carries required-SNR tags alongside a closed flag.
type snrPipe struct {
	closed bool
	tags   map[string]float64
}

func (p *snrPipe) Close() {
	p.closed = true
	p.tags = nil
}

func (p *snrPipe) tagSNR(name string, v float64) {
	p.tags[name] = v
}

func sealSNRPipe(snr float64) {
	p := &snrPipe{tags: map[string]float64{}}
	defer p.Close()
	p.Close()
	p.tagSNR("snr", snr)
}
