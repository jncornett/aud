package limit

import (
	"github.com/jncornett/aud"
)

// Source implements a transform the clips a source to a finite number of samples.
type Source struct {
	wrapped   aud.Source
	remaining int
}

// New creates a new Source from with n maximum samples.
func New(src aud.Source, limit int) *Source {
	return &Source{wrapped: src, remaining: limit}
}

// Next returns the next sample from the source.
func (src *Source) Next() (s aud.Sample, eof bool) {
	if src.remaining <= 0 {
		eof = true
		return
	}
	src.remaining--
	s, eof = src.wrapped.Next()
	return
}
