package quantize

import (
	"math"

	"github.com/jncornett/aud"
)

const (
	// MaxInt24 represents the maximum value that can be stored in a signed
	// 24-bit integer.
	MaxInt24 = 1<<23 - 1
	// MinInt24 represents the minimum value that can be stored in a signed
	// 24-bit integer.
	MinInt24 = -1 << 23
)

// To8Bit normalizes a number from the range [-1, 1] to [-127, 128] as integers.
func To8Bit(s aud.Sample) aud.Sample {
	return clip(round(s*128) + 128)
}

// To16Bit normalizes a number from the range [-1, 1] to [-32767, 32768] as integers.
func To16Bit(s aud.Sample) aud.Sample {
	return round(s * math.MaxInt16)
}

// To24Bit normalizes a number from the range [-1, 1] to [-8388607, 8388608] as integers.
func To24Bit(s aud.Sample) aud.Sample {
	return round(s * MaxInt24)
}

// To32Bit normalizes a number from the range [-1, 1] to [-2147483647, 2147483648] as integers.
func To32Bit(s aud.Sample) aud.Sample {
	return round(s * math.MaxInt32)
}

func round(s aud.Sample) aud.Sample {
	return aud.Sample(math.Floor(float64(s) + 0.5))
}

func clip(s aud.Sample) aud.Sample {
	return aud.Sample(math.Max(-1, math.Min(1, float64(s))))
}
