package function

import (
	"github.com/jncornett/aud"
)

type Gen struct {
	Func func(t int) aud.Sample
	t    int
}

func Square(off, on int, min, max aud.Sample) *Gen {
	steps := 0
	bit := false
	return &Gen{
		Func: func(t int) (s aud.Sample) {
			if bit {
				s = max
				steps++
				if steps > on {
					steps = 0
					bit = false
				}
			} else {
				s = min
				steps++
				if steps > off {
					steps = 0
					bit = true
				}
			}
			return
		},
	}
}

func (g *Gen) Next() (s aud.Sample, eof bool) {
	s = g.Func(g.t)
	g.t++
	return
}
