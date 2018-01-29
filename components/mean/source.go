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
func (m *Mixer) Next() aud.Sample {
	if len(m.sources) == 0 {
		return aud.Zero
	}
	sum := aud.Zero
	for _, src := range m.sources {
		sum += src.Next()
	}
	return sum / aud.Sample(len(m.sources))
}

func (m *Mixer) HasNext() bool {
	for _, src := range m.sources {
		if src.HasNext() {
			return true
		}
	}
	return false
}
