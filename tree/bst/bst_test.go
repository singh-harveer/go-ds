package bst_test

import (
	"golang/ds/tree/bst"
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	t.Parallel()
	var binaryTreeObj = bst.NewBinary()

	for _, value := range []int{50, 7, 9, 60, 59, 61} {
		binaryTreeObj.Insert(value)
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

			binaryTreeObj.Insert(tc.in)
			var got = binaryTreeObj.Search(tc.v)
			binaryTreeObj.LevelOrderTraversal()

			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
