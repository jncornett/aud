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

// Int8Source represents an int8 sample source.
type Int8Source interface {
	Next() (s int8, eof bool)
}

// Int16Source represents an int16 sample source.
type Int16Source interface {
	Next() (s int16, eof bool)
}

// Int32Source repesents an int32 sample source.
type Int32Source interface {
	Next() (s int32, eof bool)
}

// Int8Zipped represents bytes in an int8 channel set.
type Int8Zipped interface {
	Next() (s []int8, eof bool)
}

// Int16Zipped represents bytes in an int8 channel set.
type Int16Zipped interface {
	Next() (s []int16, eof bool)
}

// Int32Zipped represents bytes in an int8 channel set.
type Int32Zipped interface {
	Next() (s []int32, eof bool)
}

// ForEach applies a function to each sample in a Source.
func ForEach(src Source, fn func(Sample)) {
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		fn(s)
	}
}

// Max finds the maximum sample value of a Source.
func Max(src Source) Sample {
	max, eof := src.Next()
	if eof {
		return max
	}
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		if s > max {
			max = s
		}
	}
	return max
}

// Min finds the minimum sample value of a Source.
func Min(src Source) Sample {
	min, eof := src.Next()
	if eof {
		return min
	}
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		if s < min {
			min = s
		}
	}
	return min
}

// ForEachInt8 applies a function to each sample in an Int8Source.
func ForEachInt8(src Int8Source, fn func(int8)) {
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		fn(s)
	}
}

// ForEachInt16 applies a function to each sample in an Int16Source.
func ForEachInt16(src Int16Source, fn func(int16)) {
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		fn(s)
	}
}

// ForEachInt32 applies a function to each sample in an Int32Source.
func ForEachInt32(src Int32Source, fn func(int32)) {
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		fn(s)
	}
}

// ForEachInt8Slice applies a function to each sample in an Int8Zipped.
func ForEachInt8Slice(src Int8Zipped, fn func([]int8)) {
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		fn(s)
	}
}

// ForEachInt16Slice applies a function to each sample in an Int16Zipped.
func ForEachInt16Slice(src Int16Zipped, fn func([]int16)) {
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		fn(s)
	}
}

// ForEachInt32Slice applies a function to each sample in an Int32Zipped.
func ForEachInt32Slice(src Int32Zipped, fn func([]int32)) {
	for {
		s, eof := src.Next()
		if eof {
			break
		}
		fn(s)
	}
}
