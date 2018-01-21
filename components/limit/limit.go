package limit

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a transform the clips a source to a finite number of samples.
type Source struct {
	aud.Source
	Points int
}

// New creates a new Source from with n maximum samples.
func New(s aud.Source, n int) *Source {
	return &Source{s, n}
}

// Next returns the next sample from the source.
func (s *Source) Next() sample.Point {
	if s.Points <= 0 {
		return aud.EOF
	}
	s.Points--
	return s.Source.Next()
}

var _ aud.Source = new(Source)
