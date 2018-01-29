package constant

import (
	"github.com/jncornett/aud"
)

type Gen struct {
	Value aud.Sample
}

func New(v aud.Sample) *Gen {
	return &Gen{Value: v}
}

func (src *Gen) Next() (s aud.Sample, eof bool) {
	return src.Value, false
}
