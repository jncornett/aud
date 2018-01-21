package aud

// Source represents a sample source. Source is rate-agnostic, so the
// appropriate interpolation/compression should be applied to normalize all
// sources.
type Source interface {
	// Next should be implemented to return the next data point from a Source.
	// If the source is drained, eof will be true.
	Next() (s Sample, eof bool)
}

// Sample represents a sequential datapoint in the sample space.
type Sample float64

// Zero represents the zero valued sample.
const Zero Sample = 0
