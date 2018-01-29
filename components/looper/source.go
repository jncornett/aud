package looper

import (
	"github.com/jncornett/aud"
)

// Source implements a looping audio source.
type Source struct {
	factory func() aud.Source
	cur     aud.Source
}

// New creates a new Source.
func New(factory func() aud.Source) *Source {
	return &Source{factory: factory, cur: factory()}
}

// Next returns the next sample from the source.
func (src *Source) Next() (s aud.Sample) {
	if !src.cur.HasNext() {
		src.cur = src.factory()
	}
	s = src.cur.Next()
	return
}

func (src *Source) HasNext() bool {
	if !src.cur.HasNext() {
		src.cur = src.factory()
	}
	return src.cur.HasNext()
}
