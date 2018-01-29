package quantize

import (
	"math"
	"testing"

	"github.com/jncornett/aud"
	"github.com/stretchr/testify/assert"
)

func TestTo8Bit(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(aud.Sample(math.MinInt8+1), To8Bit(-1))
	assert.Equal(aud.Sample(math.MaxInt8), To8Bit(1))
}

func TestTo16Bit(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(aud.Sample(math.MinInt16+1), To16Bit(-1))
	assert.Equal(aud.Sample(math.MaxInt16), To16Bit(1))
}

func TestTo24Bit(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(aud.Sample(MinInt24+1), To24Bit(-1))
	assert.Equal(aud.Sample(MaxInt24), To24Bit(1))
}

func TestTo32Bit(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(aud.Sample(math.MinInt32+1), To32Bit(-1))
	assert.Equal(aud.Sample(math.MaxInt32), To32Bit(1))
}
