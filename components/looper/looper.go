package looper

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a looping audio source.
type Source struct {
	aud.ResettableSource
}

// New creates a new Source.
func New(rs aud.ResettableSource) *Source {
	return &Source{rs}
}

// Next returns the next sample from the source. Unless the backing source
// is empty, Next will never return aud.EOF.
func (s *Source) Next() sample.Point {
	p := s.ResettableSource.Next()
	if p == aud.EOF {
		s.Reset()
		p = s.ResettableSource.Next()
	}
	return p
}

var _ aud.ResettableSource = new(Source)
