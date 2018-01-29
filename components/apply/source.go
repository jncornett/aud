package apply

import (
	"github.com/jncornett/aud"
)

type Mapper struct {
	Func func(aud.Sample) aud.Sample
	aud.Source
}

func Map(fn func(aud.Sample) aud.Sample, src aud.Source) *Mapper {
	return &Mapper{Func: fn, Source: src}
}

func (m *Mapper) Next() (s aud.Sample, eof bool) {
	s, eof = m.Source.Next()
	s = m.Func(s)
	return
}
