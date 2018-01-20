package nlooper

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a looping audio source that loops for a finite number of
// times.
type Source struct {
	aud.ResettableSource
	N int
}

// New creates a new Source from rs and n where n is the number of times to
// loop.
func New(rs aud.ResettableSource, n int) *Source {
	return &Source{rs, n}
}

// Next returns the next sample from the source. If the source is drained,
// Next returns aud.EOF.
func (s *Source) Next() sample.Point {
	p := s.ResettableSource.Next()
	if p == aud.EOF && s.N > 0 {
		s.N--
		s.ResettableSource.Reset()
		p = s.ResettableSource.Next()
	}
	return p
}

var _ aud.ResettableSource = new(Source)
