package fifo

import (
	"github.com/jncornett/aud"
)

type Q struct {
	sources []aud.Source
}

func New(sources ...aud.Source) *Q {
	return &Q{sources}
}

func (q *Q) Peek() aud.Source {
	if len(q.sources) == 0 {
		return nil
	}
	return q.sources[0]
}

func (q *Q) Pop() {
	if len(q.sources) == 0 {
		return
	}
	q.sources = q.sources[1:]
}

func (q *Q) Len() int {
	return len(q.sources)
}
