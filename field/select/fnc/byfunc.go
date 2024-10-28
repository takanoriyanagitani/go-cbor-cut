package byfunc

import (
	sl "github.com/takanoriyanagitani/go-cbor-cut/field/select"
)

type SelectFieldsByFunc func(idx uint32) (keep bool)

func (s SelectFieldsByFunc) ToSelectFields() sl.SelectFields {
	return func(i sl.Indices) sl.Selected {
		var ret []uint32
		for _, idx := range i {
			var keep bool = s(idx)
			if keep {
				ret = append(ret, idx)
			}
		}
		return ret
	}
}
