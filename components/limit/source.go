package limit

import (
	"github.com/jncornett/aud"
)

// Source implements a transform the clips a source to a finite number of samples.
type Source struct {
	aud.Source
	remaining int
}

// New creates a new Source from with n maximum samples.
func New(src aud.Source, limit int) *Source {
	return &Source{src, limit}
}

// Next returns the next sample from the source.
func (src *Source) Next() (s aud.Sample) {
	if src.remaining <= 0 {
		return
	}
	src.remaining--
	s = src.Source.Next()
	return
}

func (src *Source) HasNext() bool {
	return src.Source.HasNext() && src.remaining > 0
}
