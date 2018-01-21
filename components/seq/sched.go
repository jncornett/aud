package seq

import (
	"github.com/jncornett/aud"
)

// Cue represents a cue point in the scheduler.
type Cue struct {
	aud.Source
	Start int
}

// SortableCue implements sort.Interface for a slice of cues.
type SortableCue []Cue

func (s SortableCue) Len() int           { return len(s) }
func (s SortableCue) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortableCue) Less(i, j int) bool { return s[i].Start < s[j].Start }
