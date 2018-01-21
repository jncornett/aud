package nlooper

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a looping audio source that loops for a finite number of
// times.
type Source struct {
	factory   func() aud.Source
	src       aud.Source
	remaining int
}

// New creates a new Source from rs and n where n is the number of times to
// loop.
func New(factory func() aud.Source, limit int) *Source {
	return &Source{
		factory:   factory,
		src:       factory(),
		remaining: limit,
	}
}

// Next returns the next sample from the source.
func (s *Source) Next() (p sample.Point, eof bool) {
	p, eof = s.src.Next()
	if eof && s.remaining > 0 {
		s.remaining--
		s.src = s.factory()
		p, eof = s.src.Next()
	}
	return
}

var _ aud.Source = new(Source)
