package sample

import (
	"math/rand"

	"github.com/jncornett/aud"
)

type Source struct {
	src  aud.Source
	rate float64
	fn   func(aud.Sample)
}

func New(src aud.Source, rate float64, fn func(aud.Sample)) *Source {
	return &Source{src, rate, fn}
}

func (src *Source) Next() (s aud.Sample, eof bool) {
	s, eof = src.src.Next()
	if rand.Float64() < src.rate {
		src.fn(s)
	}
	return
}
