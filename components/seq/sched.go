package seq

import (
	"sort"

	"github.com/jncornett/aud"
	"github.com/jncornett/aud/components/fixlen"
)

// Cue represents a cue point in the scheduler.
type Cue struct {
	aud.Source
	Start int
}

// Schedule constructs a sequential source from a set of cue points.
func Schedule(cues ...Cue) *Source {
	sched := dedupe(sorted(cues))
	sources := make([]aud.Source, 0, len(sched))
	for i, cue := range sched {
		if i >= len(sched)-1 {
			sources = append(sources, cue.Source)
			continue
		}
		next := sched[i+1].Start
		src := fixlen.New(cue.Source, next-cue.Start)
		sources = append(sources, src)
	}
	return New(sources...)
}

// sortableCue implements sort.Interface for a slice of cues.
type sortableCue []Cue

func (sc sortableCue) Len() int           { return len(sc) }
func (sc sortableCue) Swap(i, j int)      { sc[i], sc[j] = sc[j], sc[i] }
func (sc sortableCue) Less(i, j int) bool { return sc[i].Start < sc[j].Start }

func sorted(cues []Cue) []Cue {
	out := make([]Cue, len(cues))
	copy(out, cues)
	sort.Sort(sortableCue(out))
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
