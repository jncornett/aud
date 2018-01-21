package looper

import (
	"github.com/jncornett/aud"
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
func (src *Source) Next() (s aud.Sample, eof bool) {
	s, eof = src.src.Next()
	if eof {
		src.src = src.factory()
		s, eof = src.Next()
	}
	return
}
