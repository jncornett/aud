package slice

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a slice-backed sample source.k
type Source struct {
	samples []sample.Point
	pos     int
}

// New creates a new Source.
func New(samples ...sample.Point) *Source {
	return &Source{samples: samples}
}

// Next returns the next sample in the slice.
func (s *Source) Next() (p sample.Point, eof bool) {
	if s.pos >= len(s.samples) {
		eof = true
		return
	}
	p = s.samples[s.pos]
	s.pos++
	return
}

var _ aud.Source = new(Source)
