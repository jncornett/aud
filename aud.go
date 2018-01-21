package aud

import (
	"github.com/jncornett/aud/sample"
)

// Source represents a sample source. Source is rate-agnostic, so the
// appropriate interpolation/compression should be applied to normalize all
// sources.
type Source interface {
	// Next should be implemented to return the next data point from a Source.
	// If the source is drained, eof will be true.
	Next() (p sample.Point, eof bool)
}
