package byrange

import (
	"errors"
	"strconv"

	sl "github.com/takanoriyanagitani/go-cbor-cut/field/select"

	sf "github.com/takanoriyanagitani/go-cbor-cut/field/select/fnc"
)

func StringToUint32(s string) (uint32, error) {
	u, e := strconv.ParseUint(s, 10, 32)
	return uint32(u), e
}

func SelByLowerFromStr(s string) (SelectFieldsByLowerBound, error) {
	u, e := StringToUint32(s)
	return SelectFieldsByLowerBound(u), e
}

func SelByUpperFromStr(s string) (SelectFieldsByUpperBound, error) {
	u, e := StringToUint32(s)
	return SelectFieldsByUpperBound(u), e
}

func SelByRangeFromStr(lb, ub string) (SelectFieldsByRange, error) {
	l, el := StringToUint32(lb)
	u, eu := StringToUint32(ub)
	return SelectFieldsByRange{
		LowerBound: l,
		UpperBound: u,
	}, errors.Join(el, eu)
}

func SelByRangeFromStrings(s []string) (SelectFieldsByRange, error) {
	switch len(s) {
	case 2:
		return SelByRangeFromStr(s[0], s[1])
	default:
		return SelByRangeFromStr("", "") // always error
	}
}

type SelectFieldsByLowerBound uint32
type SelectFieldsByUpperBound uint32

type SelectFieldsByRange struct {
	LowerBound uint32
	UpperBound uint32
}

func (s SelectFieldsByLowerBound) ToSelByFn() sf.SelectFieldsByFunc {
	return func(idx uint32) (keep bool) {
		return uint32(s) <= idx
	}
}

func (s SelectFieldsByLowerBound) ToSelectFields() sl.SelectFields {
	return s.ToSelByFn().ToSelectFields()
}

func (s SelectFieldsByUpperBound) ToSelByFn() sf.SelectFieldsByFunc {
	return func(idx uint32) (keep bool) {
		return idx <= uint32(s)
	}
}

func (s SelectFieldsByUpperBound) ToSelectFields() sl.SelectFields {
	return s.ToSelByFn().ToSelectFields()
}

func (s SelectFieldsByRange) ToSelByFn() sf.SelectFieldsByFunc {
	return func(idx uint32) (keep bool) {
		return s.LowerBound <= idx && idx <= s.UpperBound
	}
}

func (s SelectFieldsByRange) ToSelectFields() sl.SelectFields {
	return s.ToSelByFn().ToSelectFields()
}
