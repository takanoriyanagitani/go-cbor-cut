package str2sel_test

import (
	"testing"

	sl "github.com/takanoriyanagitani/go-cbor-cut/field/select"
	ss "github.com/takanoriyanagitani/go-cbor-cut/field/select/str2sel"
)

func TestStrToSel(t *testing.T) {
	t.Parallel()

	t.Run("StringToSelectDefault", func(t *testing.T) {
		t.Parallel()

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			s, e := ss.StringToSelectDefault("")
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			var selected sl.Selected = s([]uint32{1, 2, 3, 4})
			var l int = len(selected)
			if 0 != l {
				t.Fatalf("must be empty\n")
			}
		})

		t.Run("1st", func(t *testing.T) {
			t.Parallel()

			s, e := ss.StringToSelectDefault("1")
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			var selected sl.Selected = s([]uint32{3, 1, 4})
			var l int = len(selected)
			if 1 != l {
				t.Fatalf("unexpected len: %v\n", l)
			}
			if 1 != selected[0] {
				t.Fatalf("unexpected value: %v\n", selected[0])
			}
		})

		t.Run("2nd", func(t *testing.T) {
			t.Parallel()

			s, e := ss.StringToSelectDefault("2")
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			var selected sl.Selected = s([]uint32{4, 2, 1, 9, 5})
			var l int = len(selected)
			if 1 != l {
				t.Fatalf("unexpected len: %v\n", l)
			}
			if 2 != selected[0] {
				t.Fatalf("unexpected value: %v\n", selected[0])
			}
		})

		t.Run("double", func(t *testing.T) {
			t.Parallel()

			s, e := ss.StringToSelectDefault("2,5")
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			var i []uint32 = []uint32{4, 2, 1, 9, 5}
			var selected sl.Selected = s(i)
			var l int = len(selected)
			if 2 != l {
				t.Fatalf("unexpected len: %v\n", l)
			}

			m := map[uint32]struct{}{}
			s.ToMap(i, m)
			expected := []uint32{2, 5}
			for _, ex := range expected {
				_, found := m[ex]
				var missing bool = !found
				if missing {
					t.Fatalf("value %v missing\n", ex)
				}
			}
		})

		t.Run("triple", func(t *testing.T) {
			t.Parallel()

			s, e := ss.StringToSelectDefault("1,9,5")
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			var i []uint32 = []uint32{4, 2, 1, 9, 5}
			var selected sl.Selected = s(i)
			var l int = len(selected)
			if 3 != l {
				t.Fatalf("unexpected len: %v\n", l)
			}

			m := map[uint32]struct{}{}
			s.ToMap(i, m)
			expected := []uint32{1, 9, 5}
			for _, ex := range expected {
				_, found := m[ex]
				var missing bool = !found
				if missing {
					t.Fatalf("value %v missing\n", ex)
				}
			}
		})

		t.Run("lower", func(t *testing.T) {
			t.Parallel()

			s, e := ss.StringToSelectDefault("3-")
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			var i []uint32 = []uint32{4, 2, 1, 9, 5}
			var selected sl.Selected = s(i)
			var l int = len(selected)
			if 3 != l {
				t.Fatalf("unexpected len: %v\n", l)
			}

			m := map[uint32]struct{}{}
			s.ToMap(i, m)
			expected := []uint32{4, 9, 5}
			for _, ex := range expected {
				_, found := m[ex]
				var missing bool = !found
				if missing {
					t.Fatalf("value %v missing\n", ex)
				}
			}
		})

		t.Run("upper", func(t *testing.T) {
			t.Parallel()

			s, e := ss.StringToSelectDefault("-3")
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			var i []uint32 = []uint32{4, 2, 1, 9, 5}
			var selected sl.Selected = s(i)
			var l int = len(selected)
			if 2 != l {
				t.Fatalf("unexpected len: %v\n", l)
			}

			m := map[uint32]struct{}{}
			s.ToMap(i, m)
			expected := []uint32{2, 1}
			for _, ex := range expected {
				_, found := m[ex]
				var missing bool = !found
				if missing {
					t.Fatalf("value %v missing\n", ex)
				}
			}
		})

		t.Run("range", func(t *testing.T) {
			t.Parallel()

			s, e := ss.StringToSelectDefault("2-5")
			if nil != e {
				t.Fatalf("unexpected error: %v\n", e)
			}

			var i []uint32 = []uint32{4, 2, 1, 9, 5}
			var selected sl.Selected = s(i)
			var l int = len(selected)
			if 3 != l {
				t.Fatalf("unexpected len: %v\n", l)
			}

			m := map[uint32]struct{}{}
			s.ToMap(i, m)
			expected := []uint32{4, 2, 5}
			for _, ex := range expected {
				_, found := m[ex]
				var missing bool = !found
				if missing {
					t.Fatalf("value %v missing\n", ex)
				}
			}
		})
	})
}
