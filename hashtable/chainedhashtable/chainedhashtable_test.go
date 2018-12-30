package chainedhashtable

import (
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().len; ret != 0 {
		t.Errorf("ChainedHashTable.New().len = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	cases := map[string]struct {
		cht        ChainedHashTable
		index      int
		want       int
		expectFail bool
	}{
		"add new elements": {
			cht: *New(),
			index: 10,
			want: 10,
			expectFail: false,
		},
		"add existed elements": {
			cht: *dummyChainedHashTable(10),
			index: 10,
			want: 10,
			expectFail: true,
		},
	}

	for n, tc := range cases {
		tc := tc

		t.Run(n, func (*testing.T) {
			for i := 0; i < tc.index; i++ {
				if ok := tc.cht.Add(i); !ok && !tc.expectFail {
					t.Errorf("Add returned %v unexpectedly", ok)
				}
			}

			if tc.cht.len != tc.want {
				t.Errorf("cht.len = %v but want %v", tc.cht.len, tc.want)
			}
		})
	}

}

func TestFind(t *testing.T) {
	n := 10
	cht := New()

	for i := 0; i < n; i++ {
		ret := cht.Find(i)
		if ret != nil {
			t.Errorf("Add returned non-nil unexpectedly, ret=%d", ret)
		}
	}

	for i := 0; i < n; i++ {
		cht.Add(i)
		ret := cht.Find(i)
		if ret == nil {
			t.Errorf("Add returned nil unexpectedly")
		}
	}

	ret := cht.Find(n + 123)
	if ret != nil {
		t.Errorf("Add returned nil unexpectedly")
	}
}

func TestRemove(t *testing.T) {
	x := 12345
	cht := New()

	if cht.Remove(x) == true {
		t.Errorf("Add returned true unexpectedly")
	}

	cht.Add(x)
	if cht.Remove(x) == false {
		t.Errorf("Add returned false unexpectedly")
	}

	if cht.Find(x) != nil {
		t.Errorf("Add returned non-nil unexpectedly")
	}
}

func dummyChainedHashTable(n int) *ChainedHashTable {
	cht := New()
	for i := 0; i < n; i++ {
		cht.Add(i)
	}
	return cht
}