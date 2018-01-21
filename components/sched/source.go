package sched

import (
	"sort"

	"github.com/jncornett/aud"
	"github.com/jncornett/aud/sample"
)

// Source implements a scheduling sample source.
type Source struct {
	schedule []Cue
	t        int
}

// New creates a new sample source.
func New(cues ...Cue) *Source {
	schedule := sorted(dedupe(cues))
	return &Source{schedule: schedule}
}

// Next returns the sample from the next cued source on a schedule.
func (s *Source) Next() (p sample.Point, eof bool) {
	// can the next source be cued?
	if len(s.schedule) > 1 && s.schedule[1].Start == s.t {
		s.schedule = s.schedule[1:]
	}
	if len(s.schedule) == 0 {
		// cue is empty
		return sample.Zero, true
	}
	cue := s.schedule[0]
	if cue.Start > s.t {
		// cur source is not ready yet
		s.t++
		return sample.Zero, false
	}
	p, eof = cue.Next()
	if eof {
		// current source is drained, pop it off the stack
		s.schedule = s.schedule[1:]
		if len(s.schedule) != 0 {
			// still have future scheduled events, even though current source has been drained.
			eof = false
		}
	}
	return p, eof
}

func (s *Source) cur() *Cue {
	if len(s.schedule) == 0 {
		return nil
	}
	return &s.schedule[0]
}

func (s *Source) next() *Cue {
	if len(s.schedule) <= 1 {
		return nil
	}
	return &s.schedule[1]
}

func (s *Source) pop() {
	if len(s.schedule) == 0 {
		return
	}
	s.schedule = s.schedule[1:]
}

var _ aud.Source = new(Source)

func sorted(cues []Cue) []Cue {
	out := make([]Cue, len(cues))
	copy(out, cues)
	sort.Sort(schedule(out))
	return out
}

func dedupe(cues []Cue) []Cue {
	if len(cues) == 0 {
		return nil
	}
	out := []Cue{cues[0]}
	for _, cue := range cues[1:] {
		if out[len(out)-1].Start == cue.Start {
			out[len(out)-1] = cue
			continue
		}
		out = append(out, cue)
	}
	return out
}
