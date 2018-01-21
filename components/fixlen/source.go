package fixlen

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a fixed length source that clips a source to a fixed
// length and fills in that length with sample.Zero if the underlying
// source is drained.
type Source struct {
	src       aud.Source
	remaining int
}

// New creates a new source with length.
func New(src aud.Source, length int) *Source {
	return &Source{src, length}
}

// Next returns the next sample in the sequence.
func (s *Source) Next() (p sample.Point, eof bool) {
	if s.remaining == 0 {
		eof = true
		return
	}
	s.remaining--
	p, _ = s.src.Next()
	return
}

var _ aud.Source = new(Source)
