package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jncornett/aud"
	"github.com/jncornett/aud/codecs/wav"
	"github.com/jncornett/aud/components/fixlen"
	"github.com/jncornett/aud/generators/function"
)

var (
	toneLength = flag.Duration("duration", time.Second, "tone length")
	sampleRate = flag.Uint("rate", 8000, "sample rate")
	frequency  = flag.Uint("freq", 440, "frequency (in Hertz)")
	bitDepth   = flag.Int("bits", 8, "bit depth (must be one of 8 or 16)")
)

func main() {
	flag.Parse()
	numSamples := int(*toneLength/time.Second) * int(*sampleRate)
	log.Println("Number of samples:", numSamples)
	cycle := int(aud.Hz(*sampleRate) / aud.Hz(*frequency))
	gen := function.Square(cycle/2, cycle/2, aud.Sample(-1), aud.Sample(1))
	master := fixlen.New(gen, numSamples)
	var err error
	switch *bitDepth {
	case 8:
		err = wav.EncodeMono(os.Stdout, wav.EightBitUnsigned{}, aud.Hz(*sampleRate), master)
	case 16:
		err = wav.EncodeMono(os.Stdout, wav.SixteenBitSigned{}, aud.Hz(*sampleRate), master)
	default:
		panic(fmt.Errorf("invalid bit depth: %d", *bitDepth))
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println("complete!")
}
