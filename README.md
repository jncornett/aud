# Aud

[![GoDoc](https://godoc.org/github.com/jncornett/aud?status.svg)](https://godoc.org/github.com/jncornett/aud)

A an audio mixer library implemented in pure Go.

## Design/Notes

- aud is a a compositional mixer library. Individual and simple mixer components are composed together to create more complex components.
- `aud.Source` is the basic building block of a mix. It is a single-method interface for `Next() (s Sample, eof bool)`.

## TODO

- implement resampler.
- implement mixers.
- add (empty) unit tests for all components.