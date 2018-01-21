package fixlen

import (
	"github.com/jncornett/aud"
)

// Source implements a fixed length source that clips a source to a fixed
// length and fills in that length with sample.Zero if the underlying
// source is drained.
type Source struct {
	wrapped   aud.Source
	remaining int
}

// New creates a new source with length.
func New(src aud.Source, length int) *Source {
	return &Source{wrapped: src, remaining: length}
}

// Next returns the next sample in the sequence.
func (src *Source) Next() (s aud.Sample, eof bool) {
	if src.remaining == 0 {
		eof = true
		return
	}
	src.remaining--
	s, _ = src.wrapped.Next()
	return
}
