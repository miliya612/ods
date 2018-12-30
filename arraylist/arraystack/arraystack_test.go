package arraystack

import (
	"testing"
)

func TestArrayStack_Get(t *testing.T) {
	cases := map[string]struct {
		as          ArrayStack
		index       int
		want        interface{}
		expectPanic bool
	}{
		"first element": {
			as: ArrayStack{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       0,
			want:        1,
			expectPanic: false,
		},
		"last element": {
			as: ArrayStack{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       2,
			want:        3,
			expectPanic: false,
		},
		"out of index": {
			as: ArrayStack{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       3,
			expectPanic: true,
		},
		"negative index": {
			as: ArrayStack{
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
			as := c.as
			var got interface{}

			func() {
				defer func() {
					if r := recover(); c.expectPanic && r == nil {
						t.Errorf("%s should have panicked!", n)
					}
				}()
				got = as.Get(c.index)
			}()

			if c.want != got {
				t.Errorf("want %v but got %v", c.want, got)
			}
		})
	}

}

func TestArrayStack_Set(t *testing.T) {
	cases := map[string]struct {
		as          ArrayStack
		index       int
		val         interface{}
		want        interface{}
		expectPanic bool
	}{
		"first element": {
			as: ArrayStack{
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
			as: ArrayStack{
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
			as: ArrayStack{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       3,
			val:         5,
			expectPanic: true,
		},
		"negative index": {
			as: ArrayStack{
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
			as := c.as
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
						got = as.Get(c.index)
					}
				}()
				_ = as.Set(c.index, c.val)
			}()

			if c.want != got {
				t.Errorf("want %v but got %v", c.want, got)
			}
		})
	}

}

func TestArrayStack_Add(t *testing.T) {
	cases := map[string]struct {
		as          ArrayStack
		index       int
		val         interface{}
		want        ArrayStack
		expectPanic bool
	}{
		"first position": {
			as: ArrayStack{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index: 0,
			val:   10,
			want: ArrayStack{
				buf: []interface{}{10, 1, 2, 3},
				len: 4,
				cap: 8,
			},
			expectPanic: false,
		},
		"last position": {
			as: ArrayStack{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index: 3,
			val:   5,
			want: ArrayStack{
				buf: []interface{}{1, 2, 3, 5},
				len: 4,
				cap: 8,
			},
			expectPanic: false,
		},
		"out of index": {
			as: ArrayStack{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       4,
			val:         5,
			expectPanic: true,
		},
		"negative index": {
			as: ArrayStack{
				buf: []interface{}{1, 2, 3},
				len: 3,
				cap: 3,
			},
			index:       -1,
			expectPanic: true,
		},
		//"nil array": {
		//	as: ArrayStack{
		//		buf: nil,
		//	},
		//	index: 0,
		//	val:   5,
		//	want: ArrayStack{
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
			as := c.as

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
						got = as.Get(c.index)
					}
				}()
				as.Add(c.index, c.val)
			}()

			if isArrayStackEqual(as, c.want) {
				t.Errorf("want %v but got %v", c.want, got)
			}
		})
	}

}

func isArrayStackEqual(a, b ArrayStack) bool {
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
