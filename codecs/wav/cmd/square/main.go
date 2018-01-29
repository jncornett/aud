package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/jncornett/aud"
	"github.com/jncornett/aud/codecs/wav"
	"github.com/jncornett/aud/components/fixlen"
	"github.com/jncornett/aud/components/mean"
	"github.com/jncornett/aud/components/sample"
	"github.com/jncornett/aud/generators/function"
)

var (
	toneLength = flag.Duration("duration", time.Second, "tone length")
	sampleRate = flag.Uint("rate", 8000, "sample rate")
	frequency  = flag.Uint("freq", 440, "frequency (in Hertz)")
	bitDepth   = flag.Int("bits", 8, "bit depth (must be one of 8, 16, 24 or 32)")
)

func main() {
	flag.Parse()
	numSamples := int(*toneLength/time.Second) * int(*sampleRate)
	log.Println("Number of samples:", numSamples)
	left := leftChannel(aud.Hz(*sampleRate), aud.Hz(*frequency), numSamples)
	right := rightChannel(aud.Hz(*sampleRate), aud.Hz(*frequency)*2, numSamples)
	mix := mean.New(left[0], right[0])
	var master aud.Source
	master = mix
	// master = attenuate.New(mix, 0.9999999)
	master = sample.New(master, 0.05, func(s aud.Sample) { log.Println(s) })
	err := wav.Encode(os.Stdout, *bitDepth, aud.Hz(*sampleRate), []aud.Source{master})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("complete!")
}

func leftChannel(sampleRate, frequency aud.Hz, numSamples int) []aud.Source {
	cycle := int(sampleRate / frequency)
	gen := function.Square(cycle/2, cycle/2, aud.Sample(-1), aud.Sample(1))
	src := fixlen.New(gen, numSamples)
	return []aud.Source{src}
}

func rightChannel(sampleRate, frequency aud.Hz, numSamples int) []aud.Source {
	cycle := int(sampleRate / frequency)
	gen := function.Square(cycle/2, cycle/2, aud.Sample(-1), aud.Sample(1))
	src := fixlen.New(gen, numSamples)
	return []aud.Source{src}
}
