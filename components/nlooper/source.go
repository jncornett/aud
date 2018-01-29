package nlooper

import (
	"github.com/jncornett/aud"
)

// Source implements a looping audio source that loops for a finite number of
// times.
type Source struct {
	factory   func() aud.Source
	cur       aud.Source
	remaining int
}

// New creates a new Source from rs and n where n is the number of times to
// loop.
func New(factory func() aud.Source, limit int) *Source {
	return &Source{
		factory:   factory,
		cur:       factory(),
		remaining: limit,
	}
}

// Next returns the next sample from the source.
func (src *Source) Next() (s aud.Sample) {
	if !src.cur.HasNext() && src.remaining > 0 {
		src.remaining--
		src.cur = src.factory()
	}
	s = src.cur.Next()
	return
}

func (src *Source) HasNext() bool {
	if !src.cur.HasNext() && src.remaining > 0 {
		src.remaining--
		src.cur = src.factory()
	}
	return src.cur.HasNext()
}
