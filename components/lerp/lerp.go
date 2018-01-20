package lerp

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a linear interpolation transform on a source.
type Source struct{}

// New creates a new Source with the designated from and to rates.
func New(s aud.Source, from, to sample.Rate) *Source {
	panic("not implemented")
}

// Next returns the next interpolated sample from the source. If the source is
// drained, Next returns aud.EOF.
func (s *Source) Next() sample.Point {
	panic("not implemented")
}

var _ aud.Source = new(Source)
