package mean

import (
	"github.com/jncornett/aud"
)

// Mixer implements a simple summing mixer.
type Mixer struct {
	sources []aud.Source
}

// New creates a new Source.
func New(sources ...aud.Source) *Mixer {
	return &Mixer{sources: sources}
}

// Next returns the next sample from the source.
func (src *Mixer) Next() (s aud.Sample, eof bool) {
	eofCount := 0
	for _, source := range src.sources {
		thisVal, thisEOF := source.Next()
		if thisEOF {
			eofCount++
		}
		s += thisVal
	}
	if len(src.sources) > 0 {
		eof = eofCount == len(src.sources)
		s /= aud.Sample(len(src.sources))
	}
	return
}
