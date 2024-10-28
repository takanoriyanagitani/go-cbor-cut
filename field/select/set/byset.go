package byset

import (
	sl "github.com/takanoriyanagitani/go-cbor-cut/field/select"

	sf "github.com/takanoriyanagitani/go-cbor-cut/field/select/fnc"
)

type SelectFieldsBySet map[uint32]struct{}

func (s SelectFieldsBySet) ToSelByFn() sf.SelectFieldsByFunc {
	return func(idx uint32) (keep bool) {
		_, found := s[idx]
		return found
	}
}

func (s SelectFieldsBySet) ToSelectFields() sl.SelectFields {
	return s.ToSelByFn().ToSelectFields()
}
