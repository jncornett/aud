package aud

// Source represents a sample source. Source is rate-agnostic, so the
// appropriate interpolation/compression should be applied to normalize all
// sources.
type Source interface {
	// Next should be implemented to return the next data point from a Source.
	Next() Sample
	HasNext() bool
}

// Sample represents a sequential datapoint in the sample space.
type Sample float64

// Zero represents the zero valued sample.
const Zero Sample = 0

// UInt8Source represents an uint8 sample source.
type UInt8Source interface {
	Next() uint8
	HasNext() bool
}

// Int16Source represents an int16 sample source.
type Int16Source interface {
	Next() int16
	HasNext() bool
}

// ForEach applies a function to each sample in a Source.
func ForEach(src Source, fn func(Sample)) {
	for src.HasNext() {
		fn(src.Next())
	}
}

// Max finds the maximum sample value of a Source.
func Max(src Source) Sample {
	if !src.HasNext() {
		return Zero
	}
	max := src.Next()
	for src.HasNext() {
		s := src.Next()
		if s > max {
			max = s
		}
	}
	return max
}

// Min finds the minimum sample value of a Source.
func Min(src Source) Sample {
	if !src.HasNext() {
		return Zero
	}
	min := src.Next()
	for src.HasNext() {
		s := src.Next()
		if s < min {
			min = s
		}
	}
	return min
}

// ForEachUInt8 applies a function to each sample in a UInt8Source.
func ForEachUInt8(src UInt8Source, fn func(uint8)) {
	for src.HasNext() {
		fn(src.Next())
	}
}

// ForEachUInt8Pair applies a function to a pair of UInt8Source objects.
func ForEachUInt8Pair(left, right UInt8Source, fn func(uint8, uint8)) {
	for left.HasNext() || right.HasNext() {
		fn(left.Next(), right.Next())
	}
}

// ForEachInt16 applies a function to each sample in an Int16Source.
func ForEachInt16(src Int16Source, fn func(int16)) {
	for src.HasNext() {
		fn(src.Next())
	}
}

// ForEachInt16Pair applies a function to a pair of Int16Source objects.
func ForEachInt16Pair(left, right Int16Source, fn func(int16, int16)) {
	for left.HasNext() || right.HasNext() {
		fn(left.Next(), right.Next())
	}
}
