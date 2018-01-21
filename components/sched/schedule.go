package sched

import (
	"github.com/jncornett/aud"
)

// Cue represents a cue point in the scheduler.
type Cue struct {
	aud.Source
	Start int
}

// Schedule represents a slice of cue points and implements sort.Interface
// to allow cue points to be sorted by start time.
type schedule []Cue

func (s schedule) Len() int           { return len(s) }
func (s schedule) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s schedule) Less(i, j int) bool { return s[i].Start < s[j].Start }
