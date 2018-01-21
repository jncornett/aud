package nlooper

import (
	"github.com/jncornett/aud"
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
func (src *Source) Next() (s aud.Sample, eof bool) {
	s, eof = src.src.Next()
	if eof && src.remaining > 0 {
		src.remaining--
		src.src = src.factory()
		s, eof = src.src.Next()
	}
	return
}
