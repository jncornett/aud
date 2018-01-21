package limit

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a transform the clips a source to a finite number of samples.
type Source struct {
	src       aud.Source
	remaining int
}

// New creates a new Source from with n maximum samples.
func New(src aud.Source, limit int) *Source {
	return &Source{src, limit}
}

// Next returns the next sample from the source.
func (s *Source) Next() (p sample.Point, eof bool) {
	if s.remaining <= 0 {
		eof = true
		return
	}
	s.remaining--
	p, eof = s.src.Next()
	return
}

var _ aud.Source = new(Source)
