package sel

import (
	"iter"
	"maps"
	"slices"

	cc "github.com/takanoriyanagitani/go-cbor-cut/cut"
)

type Indices []uint32
type Selected []uint32

type SelectFields func(Indices) Selected

var SelectFieldsAll SelectFields = func(i Indices) Selected {
	return Selected(i)
}

var SelectFieldsNone SelectFields = func(_ Indices) Selected { return nil }

func (f SelectFields) ToMap(i Indices, m map[uint32]struct{}) {
	var s Selected = f(i)
	for _, idx := range s {
		m[idx] = struct{}{}
	}
}

func (f SelectFields) ToArrayToSelected() cc.ArrayToSelected {
	return func(o cc.Original) cc.Selected {
		var indices []uint32 = make([]uint32, 0, len(o))
		for i := range o {
			indices = append(indices, uint32(i))
		}
		m := map[uint32]struct{}{}
		f.ToMap(indices, m)
		var ret []any = make([]any, 0, len(m))
		for idx, val := range o {
			_, found := m[uint32(idx)]
			if found {
				ret = append(ret, val)
			}
		}
		return ret
	}
}

type SelectFieldsArray []SelectFields

func (a SelectFieldsArray) ToSelectFields() SelectFields {
	return func(i Indices) Selected {
		selected := map[uint32]struct{}{}
		for _, sf := range a {
			sf.ToMap(i, selected)
		}
		var ikeys iter.Seq[uint32] = maps.Keys(selected)
		var keys []uint32 = slices.Collect(ikeys)
		return Selected(keys)
	}
}
