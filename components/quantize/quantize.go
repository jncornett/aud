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

// To8BitUnsigned normalizes a number from the range [-1, 1] to [-127, 128] as integers.
func To8BitUnsigned(s aud.Sample) aud.Sample {
	if s <= 0 {
		return round(clip(s)*128 + 128)
	}
	return round(clip(s)*128 + 127)
}

// To16BitSigned normalizes a number from the range [-1, 1] to [-32767, 32768] as integers.
func To16BitSigned(s aud.Sample) aud.Sample {
	if s <= 0 {
		return round(clip(s) * 32768)
	}
	return round(clip(s) * 32767)
}

// To24BitSigned normalizes a number from the range [-1, 1] to [-8388607, 8388608] as integers.
func To24BitSigned(s aud.Sample) aud.Sample {
	if s <= 0 {
		return round(clip(s) * 16777216)
	}
	return round(clip(s) * 16777215)
}

func round(s aud.Sample) aud.Sample {
	return aud.Sample(math.Floor(float64(s) + 0.5))
}

func clip(s aud.Sample) aud.Sample {
	return aud.Sample(math.Max(-1, math.Min(1, float64(s))))
}
