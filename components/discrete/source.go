package discrete

import (
	"github.com/jncornett/aud"
)

type Casted8Bit struct {
	aud.Source
}

func Cast8Bit(src aud.Source) *Casted8Bit {
	return &Casted8Bit{src}
}

func (c *Casted8Bit) Next() (u uint8, eof bool) {
	s, eof := c.Source.Next()
	return uint8(s), eof
}

type Casted16Bit struct {
	aud.Source
}

func Cast16Bit(src aud.Source) *Casted16Bit {
	return &Casted16Bit{src}
}

func (c *Casted16Bit) Next() (i int16, eof bool) {
	s, eof := c.Source.Next()
	return int16(s), eof
}
