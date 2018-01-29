package attenuate

import (
	"github.com/jncornett/aud"
)

type Source struct {
	src  aud.Source
	Rate float64
}

func New(src aud.Source, rate float64) *Source {
	return &Source{src, rate}
}

func (src *Source) Next() (s aud.Sample, eof bool) {
	s, eof = src.src.Next()
	s *= aud.Sample(src.Rate)
	return
}
