package sched

import (
	"sort"

	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a scheduling sample source.
type Source struct {
	schedule Schedule
	t        int
}

// New creates a new sample source.
func New(cues ...Cue) *Source {
	sched := make(Schedule, 0, len(cues))
	copy(sched, cues)
	sort.Sort(sched)
	return &Source{schedule: sched}
}

// Next returns the sample from the next cued source on a schedule.
func (s *Source) Next() (p sample.Point, eof bool) {
	panic("not implemented")
}

var _ aud.Source = new(Source)
