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

func (g *Gen) Next() aud.Sample {
	return g.Value
}

func (g *Gen) HasNext() bool {
	return true
}
