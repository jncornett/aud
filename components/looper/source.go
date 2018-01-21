package looper

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a looping audio source.
type Source struct {
	factory func() aud.Source
	src     aud.Source
}

// New creates a new Source.
func New(factory func() aud.Source) *Source {
	return &Source{factory: factory, src: factory()}
}

// Next returns the next sample from the source.
func (s *Source) Next() (p sample.Point, eof bool) {
	p, eof = s.src.Next()
	if eof {
		s.src = s.factory()
		p, eof = s.Next()
	}
	return
}

var _ aud.Source = new(Source)
