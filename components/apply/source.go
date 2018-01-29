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
