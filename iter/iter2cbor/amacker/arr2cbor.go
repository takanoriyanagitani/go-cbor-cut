package arr2cbor

import (
	"context"
	"io"

	ca "github.com/fxamacker/cbor/v2"

	ic "github.com/takanoriyanagitani/go-cbor-cut/iter/iter2cbor"
)

type ArrToCbor struct {
	*ca.Encoder
}

func (a ArrToCbor) EncodeArray(arr []any) error {
	return a.Encoder.Encode(arr)
}

func (a ArrToCbor) ToArrayToCbor() ic.ArrayToCbor {
	return func(_ context.Context, arr []any) error {
		return a.EncodeArray(arr)
	}
}

func ArrToCborNew(w io.Writer) ArrToCbor {
	return ArrToCbor{
		Encoder: ca.NewEncoder(w),
	}
}
