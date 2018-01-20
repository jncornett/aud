package aud

import "github.com/jncornett/aud/sample"

// EOF signifies that a source is drained.
const EOF = sample.Missing

// Source represents a sample source. Source is rate-agnostic, so the
// appropriate interpolation/compression should be applied to normalize all
// sources.
type Source interface {
	// Next should be implemented to return the next data point from a Source.
	// If the source is drained, Next should return EOF.
	Next() sample.Point
}

// Resetter represents an object that implements the Reset function.
type Resetter interface {
	Reset()
}

// ResettableSource represents an object that implements the Source and
// Resetter interfaces.
type ResettableSource interface {
	Source
	Resetter
}
