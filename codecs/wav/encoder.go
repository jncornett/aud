package wav

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/jncornett/aud/components/apply"
	"github.com/jncornett/aud/components/discrete"
	"github.com/jncornett/aud/components/quantize"
	"github.com/jncornett/aud/zip"
	"github.com/jncornett/unpanic"

	"github.com/jncornett/aud"
)

const (
	headerFmtChunkSize = uint32(16)
	fileHeaderSize     = uint32(12 + headerFmtChunkSize) // waveFourCC + fmtFourCC + fmtChunkSize + remainingHeaderSize
	pmcWavFormat       = uint16(1)
)

var (
	riffFourCC                 = [4]byte{'R', 'I', 'F', 'F'}
	waveFourCC                 = [4]byte{'W', 'A', 'V', 'E'}
	fmtFourCC                  = [4]byte{'f', 'm', 't', ' '}
	dataFourCC                 = [4]byte{'d', 'a', 't', 'a'}
	errInconsistentNumChannels = errors.New("chunks have an inconsistent number of channels")
)

// Encode ...
func Encode(w io.Writer, bitDepth int, sampleRate aud.Hz, chunks ...[]aud.Source) (err error) {
	defer unpanic.Handle(&err)
	nChannels := getNumChannels(chunks...)
	// validate bit depth
	switch bitDepth {
	case 8, 16, 24, 32:
	default:
		panic(fmt.Errorf("invalid bit depth: %d", bitDepth))
	}
	var (
		dataBuf  bytes.Buffer
		chunkBuf bytes.Buffer
	)
	for _, chunk := range chunks {
		chunkBuf.Reset()
		writeChunkData(&chunkBuf, bitDepth, chunk)
		writeLE(&dataBuf, dataFourCC)
		writeLE(&dataBuf, uint32(chunkBuf.Len()))
		copy(&dataBuf, &chunkBuf)
	}
	writeFileHeader(w, nChannels, bitDepth, int(sampleRate), dataBuf.Len())
	copy(w, &dataBuf)
	return
}

func writeFileHeader(w io.Writer, numChannels, bitDepth, sampleRate, fileContentSize int) {
	blockAlign := numChannels * bitDepth / 8
	writeLE(w, riffFourCC)
	writeLE(w, fileHeaderSize+uint32(fileContentSize)) // size of this header + contents
	writeLE(w, waveFourCC)
	writeLE(w, fmtFourCC)
	writeLE(w, headerFmtChunkSize)
	writeLE(w, pmcWavFormat)
	writeLE(w, uint16(numChannels))
	writeLE(w, uint32(sampleRate))
	writeLE(w, uint32(sampleRate*blockAlign)) // avg bytes/sec
	writeLE(w, uint16(blockAlign))
	writeLE(w, uint16(bitDepth))
}

func writeChunkData(w io.Writer, bitDepth int, chunk []aud.Source) {
	quantized := make([]aud.Source, 0, len(chunk))
	// warning: code duplication ahead
	switch bitDepth {
	case 8:
		for _, src := range chunk {
			quantized = append(quantized, apply.Map(quantize.To8Bit, src))
		}
		casted := make([]aud.Int8Source, 0, len(chunk))
		for _, src := range quantized {
			casted = append(casted, discrete.Cast8Bit(src))
		}
		zipped := zip.NewInt8Zipper(casted...)
		aud.ForEachInt8Slice(zipped, func(v []int8) { writeLE(w, v) })
	case 16:
		for _, src := range chunk {
			quantized = append(quantized, apply.Map(quantize.To16Bit, src))
		}
		casted := make([]aud.Int16Source, 0, len(chunk))
		for _, src := range quantized {
			casted = append(casted, discrete.Cast16Bit(src))
		}
		zipped := zip.NewInt16Zipper(casted...)
		aud.ForEachInt16Slice(zipped, func(v []int16) { writeLE(w, v) })
	case 24:
		for _, src := range chunk {
			quantized = append(quantized, apply.Map(quantize.To24Bit, src))
		}
		casted := make([]aud.Int32Source, 0, len(chunk))
		for _, src := range quantized {
			casted = append(casted, discrete.Cast32Bit(src))
		}
		zipped := zip.NewInt32Zipper(casted...)
		aud.ForEachInt32Slice(zipped, func(v []int32) { writeLE(w, v) })
	case 32:
		for _, src := range chunk {
			quantized = append(quantized, apply.Map(quantize.To32Bit, src))
		}
		casted := make([]aud.Int32Source, 0, len(chunk))
		for _, src := range quantized {
			casted = append(casted, discrete.Cast32Bit(src))
		}
		zipped := zip.NewInt32Zipper(casted...)
		aud.ForEachInt32Slice(zipped, func(v []int32) { writeLE(w, v) })
	default:
		// should not be reachable since we do the validation up-front in Encode
		panic(fmt.Errorf("invalid bit depth: %d", bitDepth))
	}
}

func copy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}

func writeLE(w io.Writer, v interface{}) {
	err := binary.Write(w, binary.LittleEndian, v)
	if err != nil {
		panic(err)
	}
}

func getNumChannels(chunks ...[]aud.Source) int {
	if len(chunks) == 0 {
		return 0
	}
	first := len(chunks[0])
	for _, ck := range chunks[1:] {
		if first != len(ck) {
			panic(errInconsistentNumChannels)
		}
	}
	return first
}
