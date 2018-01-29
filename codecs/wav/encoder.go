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

type Stereo struct {
	Left, Right aud.Source
}

type EightBitUnsigned struct{}

func (EightBitUnsigned) encodeChunkDataMono(w io.Writer, src aud.Source) {
	quantized := apply.Map(quantize.To8BitUnsigned, src)
	casted := discrete.Cast8Bit(quantized)
	aud.ForEachUInt8(casted, func(v uint8) { writeLE(w, v) })
}

func (EightBitUnsigned) encodeChunkDataStereo(w io.Writer, left, right aud.Source) {
	ql := apply.Map(quantize.To8BitUnsigned, left)
	cl := discrete.Cast8Bit(ql)
	qr := apply.Map(quantize.To8BitUnsigned, right)
	cr := discrete.Cast8Bit(qr)
	aud.ForEachUInt8Pair(cl, cr, func(l, r uint8) {
		writeLE(w, l)
		writeLE(w, r)
	})
}

type SixteenBitSigned struct{}

func (SixteenBitSigned) encodeChunkDataMono(w io.Writer, src aud.Source) {
	quantized := apply.Map(quantize.To16BitSigned, src)
	casted := discrete.Cast16Bit(quantized)
	aud.ForEachInt16(casted, func(v int16) { writeLE(w, v) })
}

func (SixteenBitSigned) encodeChunkDataStereo(w io.Writer, left, right aud.Source) {
	ql := apply.Map(quantize.To16BitSigned, left)
	cl := discrete.Cast16Bit(ql)
	qr := apply.Map(quantize.To16BitSigned, right)
	cr := discrete.Cast16Bit(qr)
	aud.ForEachInt16Pair(cl, cr, func(l, r int16) {
		writeLE(w, l)
		writeLE(w, r)
	})
}

func EncodeMono(w io.Writer, bitDepth interface{}, sampleRate aud.Hz, chunks ...aud.Source) (err error) {
	defer unpanic.Handle(&err)
	var (
		encodeChunk func(io.Writer, aud.Source)
		numBits     int
	)
	switch t := bitDepth.(type) {
	case EightBitUnsigned:
		encodeChunk = t.encodeChunkDataMono
		numBits = 8
	case SixteenBitSigned:
		encodeChunk = t.encodeChunkDataMono
		numBits = 16
	default:
		panic(fmt.Errorf("unknown bit depth: %T", bitDepth))
	}
	var (
		dataBuf  bytes.Buffer
		chunkBuf bytes.Buffer
	)
	for _, chunk := range chunks {
		chunkBuf.Reset()
		encodeChunk(&chunkBuf, chunk)
		writeLE(&dataBuf, dataFourCC)
		writeLE(&dataBuf, uint32(chunkBuf.Len()))
		copy(&dataBuf, &chunkBuf)
	}
	writeFileHeader(w, 1, numBits, int(sampleRate), dataBuf.Len())
	copy(w, &dataBuf)
	return
}

func EncodeStereo(w io.Writer, bitDepth interface{}, sampleRate aud.Hz, chunks ...Stereo) (err error) {
	defer unpanic.Handle(&err)
	var (
		encodeChunk func(io.Writer, aud.Source, aud.Source)
		numBits     int
	)
	switch t := bitDepth.(type) {
	case EightBitUnsigned:
		encodeChunk = t.encodeChunkDataStereo
		numBits = 8
	case SixteenBitSigned:
		encodeChunk = t.encodeChunkDataStereo
		numBits = 16
	default:
		panic(fmt.Errorf("unknown bit depth: %T", bitDepth))
	}
	var (
		dataBuf  bytes.Buffer
		chunkBuf bytes.Buffer
	)
	for _, chunk := range chunks {
		chunkBuf.Reset()
		encodeChunk(&chunkBuf, chunk.Left, chunk.Right)
		writeLE(&dataBuf, dataFourCC)
		writeLE(&dataBuf, uint32(chunkBuf.Len()))
		copy(&dataBuf, &chunkBuf)
	}
	writeFileHeader(w, 2, numBits, int(sampleRate), dataBuf.Len())
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
