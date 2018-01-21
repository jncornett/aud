package logrange

import (
	"math"

	"github.com/jncornett/aud"
)

type Compressor struct {
	src aud.Source
}

func New(src aud.Source) *Compressor {
	return &Compressor{src: src}
}

func (c *Compressor) Next() (s aud.Sample, eof bool) {
	s, eof = c.src.Next()
	if eof {
		return
	}
	if s < -1 {
		s = aud.Sample(-math.Log(-float64(s)-0.85)/14 - 0.75)
		return
	}
	if s > 1 {
		s = aud.Sample(math.Log(float64(s)-0.85)/14 + 0.75)
		return
	}
	s = aud.Sample(s / 1.61803398875)
	return
}
