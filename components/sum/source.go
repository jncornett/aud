package sum

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
	eof = true
	for _, source := range src.sources {
		thisVal, thisEOF := source.Next()
		if !thisEOF {
			// signifies that at least one source is not empty.
			eof = false
		}
		s += thisVal
	}
	return
}
