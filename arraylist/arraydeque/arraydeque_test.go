package arraydeque

import (
	"testing"
)

func TestArrayDeque_Get(t *testing.T) {
	cases := map[string]struct {
		ad          ArrayDeque
		index       int
		want        interface{}
		expectPanic bool
	}{
		"first element": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       0,
			want:        1,
			expectPanic: false,
		},
		"last element": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       2,
			want:        3,
			expectPanic: false,
		},
		"out of index": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       5,
			expectPanic: true,
		},
		"negative index": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       -1,
			expectPanic: true,
		},
		"nil array": {
			expectPanic: true,
		},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			ad := c.ad
			var got interface{}

			func() {
				defer func() {
					if r := recover(); c.expectPanic && r == nil {
						t.Errorf("%s should have panicked!", n)
					}
				}()
				got = ad.Get(c.index)
			}()

			if c.want != got {
				t.Errorf("want %v but got %v", c.want, got)
			}
		})
	}

}

func TestArrayDeque_Set(t *testing.T) {
	cases := map[string]struct {
		ad          ArrayDeque
		index       int
		val         interface{}
		want        interface{}
		expectPanic bool
	}{
		"first element": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       0,
			val:         10,
			want:        10,
			expectPanic: false,
		},
		"last element": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       2,
			val:         5,
			want:        5,
			expectPanic: false,
		},
		"out of index": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       5,
			val:         5,
			expectPanic: true,
		},
		"negative index": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       -1,
			expectPanic: true,
		},
		"nil array": {
			index:       0,
			val:         5,
			expectPanic: true,
		},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			ad := c.ad
			var got interface{}
			func() {
				defer func() {
					r := recover()
					if c.expectPanic && r == nil {
						t.Errorf("%s should have panicked!", n)
					}
					if !c.expectPanic && r != nil {
						t.Errorf("%s should not have panicked!", n)
					}
					if !c.expectPanic && r == nil {
						got = ad.Get(c.index)
					}
				}()
				_ = ad.Set(c.index, c.val)
			}()

			if c.want != got {
				t.Errorf("want %v but got %v", c.want, got)
			}
		})
	}

}

func TestArrayDeque_Add(t *testing.T) {
	cases := map[string]struct {
		ad          ArrayDeque
		index       int
		val         interface{}
		want        ArrayDeque
		expectPanic bool
	}{
		"first position": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index: 0,
			val:   10,
			want: ArrayDeque{
				buf: []interface{}{10, 1, 2, 3},
				len: 4,
				cap: 8,
			},
			expectPanic: false,
		},
		"ladt position": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index: 3,
			val:   5,
			want: ArrayDeque{
				buf: []interface{}{1, 2, 3, 5},
				len: 4,
				cap: 8,
			},
			expectPanic: false,
		},
		"negative index": {
			ad: ArrayDeque{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       -1,
			expectPanic: true,
		},
		//"nil array": {
		//	ad: ArrayDeque{
		//		buf: nil,
		//	},
		//	index: 0,
		//	val:   5,
		//	want: ArrayDeque{
		//		buf: []interface{}{5},
		//		len: 1,
		//		cap: 1,
		//	},
		//	expectPanic: false,
		//},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			ad := c.ad

			var got interface{}
			func() {
				defer func() {
					r := recover()
					if c.expectPanic && r == nil {
						t.Errorf("%s should have panicked!", n)
					}
					if !c.expectPanic && r != nil {
						t.Errorf("%s should not have panicked!", n)
					}
					if !c.expectPanic && r == nil {
						got = ad.Get(c.index)
					}
				}()
				ad.Add(c.index, c.val)
			}()

			if isArrayDequeEqual(ad, c.want) {
				t.Errorf("want %v but got %v", c.want, got)
			}
		})
	}

}

func isArrayDequeEqual(a, b ArrayDeque) bool {
	if a.cap != b.cap {
		return false
	}
	if a.len != b.len {
		return false
	}
	return isSliceEqual(a.buf, b.buf)
}

func isSliceEqual(a, b []interface{}) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
