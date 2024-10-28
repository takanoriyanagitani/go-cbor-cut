package cut

import (
	"context"
	"iter"

	ct "github.com/takanoriyanagitani/go-cbor-cut"
)

type Original []any
type Selected []any

type ArrayToSelected func(Original) Selected

type ArraySource func(context.Context) iter.Seq[[]any]
type ArrayOutput func(context.Context, []any) error

type OutputSelected struct {
	ArrayToSelected
	ArraySource
	ArrayOutput
}

func (o OutputSelected) ConvertAll(ctx context.Context) error {
	var i iter.Seq[[]any] = o.ArraySource(ctx)
	for arr := range i {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		var selected Selected = o.ArrayToSelected(arr)
		e := o.ArrayOutput(ctx, selected)
		if nil != e {
			return e
		}
	}
	return nil
}

func (o OutputSelected) AsCut() ct.Cut {
	return o.ConvertAll
}
