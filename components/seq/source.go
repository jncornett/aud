package seq

import (
	"github.com/jncornett/aud"
	"github.com/jncornett/aud/fifo"
	"github.com/jncornett/aud/sample"
)

// Source concatenates a sequence of sources into a single source.
type Source struct {
	sources *fifo.Q
}

// New creates a new Source from sources.
func New(sources ...aud.Source) *Source {
	return &Source{sources: fifo.New(sources...)}
}

// Next returns the next sample from the source.
func (s *Source) Next() (p sample.Point, eof bool) {
	for s.sources.Len() > 0 {
		p, eof = s.sources.Peek().Next()
		if !eof {
			break
		}
		s.sources.Pop()
	}
	return
}

var _ aud.Source = new(Source)
