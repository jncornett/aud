package seq

import (
	"github.com/jncornett/aud"
)

// Source concatenates a sequence of sources into a single source.
type Source struct {
	sources []aud.Source
}

// New creates a new Source from sources.
func New(sources ...aud.Source) *Source {
	return &Source{sources: sources}
}

// Next returns the next sample from the source.
func (src *Source) Next() (s aud.Sample) {
	for len(src.sources) > 0 {
		cur := src.sources[0]
		if cur.HasNext() {
			s = cur.Next()
			return
		}
		src.sources = src.sources[1:]
	}
	return
}

func (src *Source) HasNext() bool {
	return len(src.sources) > 0
}
