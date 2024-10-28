package iter2cbor

import (
	"context"

	ct "github.com/takanoriyanagitani/go-cbor-cut/cut"
)

type ArrayToCbor func(context.Context, []any) error

func (a ArrayToCbor) AsArrayOutput() ct.ArrayOutput {
	return ct.ArrayOutput(a)
}
