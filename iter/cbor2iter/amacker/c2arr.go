package cbor2iter

import (
	"io"
	"iter"

	ca "github.com/fxamacker/cbor/v2"

	ci "github.com/takanoriyanagitani/go-cbor-cut/iter/cbor2iter"
)

type CborToArrays struct {
	*ca.Decoder
}

func (c CborToArrays) ToArrays() iter.Seq[[]any] {
	return func(yield func([]any) bool) {
		var buf []any
		var err error
		for {
			clear(buf)
			buf = buf[:0]
			err = c.Decoder.Decode(&buf)
			if nil != err {
				return
			}
			if !yield(buf) {
				return
			}
		}
	}
}

func (c CborToArrays) AsCborToArrays() ci.CborToArrays {
	return c.ToArrays
}

func CborToArraysNew(rdr io.Reader) CborToArrays {
	return CborToArrays{
		Decoder: ca.NewDecoder(rdr),
	}
}
