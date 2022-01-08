package tree_test

import (
	"golang/ds/tree"
	"math/rand"
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	t.Parallel()
	var bst = tree.NewBinary()

	for i := 0; i < 20; i++ {
		var value = rand.Intn(500)
		bst.Insert(value)
	}

	var testcases = []struct {
		name string
		in   int
		v    int
		want bool
	}{
		{
			name: "Ok",
			in:   501,
			v:    501,
			want: true,
		},
		{
			name: "notFound",
			in:   8,
			v:    10,
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			bst.Insert(tc.in)
			var got = bst.Search(tc.v)

			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
