package sched

import (
	"testing"

	"github.com/jncornett/aud"
	"github.com/jncornett/aud/components/slice"
	"github.com/jncornett/aud/sample"
	"github.com/stretchr/testify/assert"
)

func drain(s aud.Source) []sample.Point {
	var out []sample.Point
	for {
		p, eof := s.Next()
		if eof {
			break
		}
		out = append(out, p)
	}
	return out
}

func TestSource(t *testing.T) {
	tests := []struct {
		name   string
		cues   []Cue
		output []sample.Point
	}{
		{name: "empty schedule"},
		{
			name: "one element, one cue, no offset",
			cues: []Cue{
				Cue{
					Source: slice.New(sample.Point(1)),
					Start:  0,
				},
			},
			output: []sample.Point{sample.Point(1)},
		},
		// TODO add test cases for schedules with multiple values
		// TODO add test cases for schedules with gaps and without gaps
		// TODO add test cases for schedules with overlaps
		// TODO add test cases for schedules with negative start times
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.cues...)
			assert.Equal(t, tt.output, drain(s))
		})
	}
}
