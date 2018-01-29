package fixlen

import (
	"github.com/jncornett/aud"
)

// Source implements a fixed length source that clips a source to a fixed
// length and fills in that length with sample.Zero if the underlying
// source is drained.
type Source struct {
	aud.Source
	remaining int
}

// New creates a new source with length.
func New(src aud.Source, length int) *Source {
	return &Source{src, length}
}

// Next returns the next sample in the sequence.
func (src *Source) Next() (s aud.Sample) {
	if src.remaining <= 0 {
		return
	}
	src.remaining--
	s = src.Source.Next()
	return
}

func (src *Source) HasNext() bool {
	return src.remaining > 0
}
