package slice

import (
	"github.com/jncornett/aud"
)

// Source implements a slice-backed sample source.
type Source struct {
	samples []aud.Sample
	pos     int
}

// New creates a new Source.
func New(samples ...aud.Sample) *Source {
	return &Source{samples: samples}
}

// Next returns the next sample in the slice.
func (src *Source) Next() (s aud.Sample, eof bool) {
	if src.pos >= len(src.samples) {
		eof = true
		return
	}
	s = src.samples[src.pos]
	src.pos++
	return
}
