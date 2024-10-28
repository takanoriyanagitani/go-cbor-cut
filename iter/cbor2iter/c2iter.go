package cbor2iter

import (
	"context"
	"iter"

	ct "github.com/takanoriyanagitani/go-cbor-cut/cut"
)

type CborToArrays func() iter.Seq[[]any]

func (c CborToArrays) ToArraySource() ct.ArraySource {
	return func(_ context.Context) iter.Seq[[]any] {
		return c()
	}
}
