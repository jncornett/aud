# Aud

[![GoDoc](https://godoc.org/github.com/jncornett/aud?status.svg)](https://godoc.org/github.com/jncornett/aud)

A an audio mixer library implemented in pure Go.

## Design/Notes

- Consider changing interface of aud.Source to
  ```golang
  type Source interface {
      Next() sample.Point
      EOF() bool
  }
  ```
  This would allow `sample.Point` to represent the entire range of real values at the cost of breaking the nice single-method property of `aud.Source`.
  Alternatively, we could go the route of `io.Reader` and have multiple return values:
  ```golang
  type Source interface {
      Next() (p sample.Point, eof bool)
  }

- TODO implement lerp source.
- TODO flesh out all unit tests.