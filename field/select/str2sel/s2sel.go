package str2sel

import (
	"strconv"
	"strings"

	sl "github.com/takanoriyanagitani/go-cbor-cut/field/select"

	sr "github.com/takanoriyanagitani/go-cbor-cut/field/select/range"
	ss "github.com/takanoriyanagitani/go-cbor-cut/field/select/set"
)

const (
	DelimiterDefault string = ","
)

func StrToSel(s string, m map[uint32]struct{}) (sl.SelectFields, error) {
	before, lower := strings.CutSuffix(s, "-")
	after, upper := strings.CutPrefix(s, "-")
	var open bool = lower || upper
	var closed bool = !open
	var rng bool = closed && strings.Contains(s, "-")
	var set bool = !strings.Contains(s, "-")

	var splited []string = strings.SplitN(s, "-", 2)

	lsel, le := sr.SelByLowerFromStr(before)
	usel, ue := sr.SelByUpperFromStr(after)
	rsel, re := sr.SelByRangeFromStrings(splited)

	idx, se := strconv.ParseUint(s, 10, 32)

	switch {
	case set && nil == se:
		m[uint32(idx)] = struct{}{}
		return ss.SelectFieldsBySet(m).ToSelectFields(), nil
	case lower:
		return lsel.ToSelectFields(), le
	case upper:
		return usel.ToSelectFields(), ue
	case rng:
		return rsel.ToSelectFields(), re
	}
	return sl.SelectFieldsNone, nil
}

func StringsToSelects(all []string) ([]sl.SelectFields, error) {
	var ret []sl.SelectFields
	m := map[uint32]struct{}{}
	for _, s := range all {
		sf, e := StrToSel(s, m)
		if nil != e {
			return nil, e
		}
		ret = append(ret, sf)
	}
	return ret, nil // TODO
}

func StringsToSelect(all []string) (sl.SelectFields, error) {
	arr, e := StringsToSelects(all)
	return sl.SelectFieldsArray(arr).ToSelectFields(), e
}

func StringToSelect(all string, delim string) (sl.SelectFields, error) {
	var splited []string = strings.Split(all, delim)
	return StringsToSelect(splited)
}

func StringToSelectDefault(all string) (sl.SelectFields, error) {
	return StringToSelect(all, DelimiterDefault)
}
