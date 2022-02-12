package combinations

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPermutationsOfChar(t *testing.T) {
	var testcases = []struct {
		name string
		in   []string
		want [][]string
	}{
		{
			name: "Ok",
			in:   []string{"A", "B", "C"},
			want: [][]string{
				{"A", "B", "C"},
				{"B", "A", "C"},
				{"B", "C", "A"},
				{"A", "C", "B"},
				{"C", "A", "B"},
				{"C", "B", "A"},
			},
		},
		{
			name: "emty-slice",
			in:   []string{},
			want: [][]string{{}},
		},
	}

	for _, tc := range testcases {
		var tc = tc
		t.Run(tc.name, func(t *testing.T) {
			var got = permutationOfChar(tc.in)
			if !cmp.Equal(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
