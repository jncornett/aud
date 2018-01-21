package seq

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source concatenates a sequence of sources into a single source.
type Source struct {
	Sequence []aud.Source
}

// New creates a new Source from sources.
func New(sources ...aud.Source) *Source {
	return &Source{sources}
}

// Next returns the next sample from the source.
func (s *Source) Next() sample.Point {
	for _, x := range s.Sequence {
		p := x.Next()
		if p == aud.EOF {
			// FIXME possibly memory inefficient -- the underlying array will
			// not be garbage collected until the sequence is exhausted.
			s.Sequence = s.Sequence[1:]
		}
		return p
	}
	return aud.EOF
}

var _ aud.Source = new(Source)
