package zip

import "github.com/jncornett/aud"

// Int8Zipper zips channels of 8-bit sources.
type Int8Zipper struct {
	Sources []aud.Int8Source
	buffer  []int8
}

// NewInt8Zipper creates a new Int8Zipper.
func NewInt8Zipper(sources ...aud.Int8Source) *Int8Zipper {
	return &Int8Zipper{
		Sources: sources,
		buffer:  make([]int8, len(sources)),
	}
}

// Next returns the next values for all sources. If all sources return EOF,
// then Next will return EOF.
func (z *Int8Zipper) Next() ([]int8, bool) {
	allEOF := true
	for i, src := range z.Sources {
		s, eof := src.Next()
		z.buffer[i] = s
		if !eof {
			allEOF = false
		}
	}
	return z.buffer, allEOF
}

// Buffer returns the persistent buffer used to store results from Next.
func (z *Int8Zipper) Buffer() []int8 {
	return z.buffer
}

// Int16Zipper zips channels of 16-bit sources.
type Int16Zipper struct {
	Sources []aud.Int16Source
	buffer  []int16
}

// NewInt16Zipper creates a new Int8Zipper.
func NewInt16Zipper(sources ...aud.Int16Source) *Int16Zipper {
	return &Int16Zipper{
		Sources: sources,
		buffer:  make([]int16, len(sources)),
	}
}

// Next returns the next values for all sources. If all sources return EOF,
// then Next will return EOF.
func (z *Int16Zipper) Next() ([]int16, bool) {
	allEOF := true
	for i, src := range z.Sources {
		s, eof := src.Next()
		z.buffer[i] = s
		if !eof {
			allEOF = false
		}
	}
	return z.buffer, allEOF
}

// Buffer returns the persistent buffer used to store results from Next.
func (z *Int16Zipper) Buffer() []int16 {
	return z.buffer
}

// Int32Zipper zips channels of 32-bit sources.
type Int32Zipper struct {
	Sources []aud.Int32Source
	buffer  []int32
}

// NewInt32Zipper creates a new Int32Zipper.
func NewInt32Zipper(sources ...aud.Int32Source) *Int32Zipper {
	return &Int32Zipper{
		Sources: sources,
		buffer:  make([]int32, len(sources)),
	}
}

// Next returns the next values for all sources. If all sources return EOF,
// then Next will return EOF.
func (z *Int32Zipper) Next() ([]int32, bool) {
	allEOF := true
	for i, src := range z.Sources {
		s, eof := src.Next()
		z.buffer[i] = s
		if !eof {
			allEOF = false
		}
	}
	return z.buffer, allEOF
}

// Buffer returns the persistent buffer used to store results from Next.
func (z *Int32Zipper) Buffer() []int32 {
	return z.buffer
}
