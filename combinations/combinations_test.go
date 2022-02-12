package combinations

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCombinationOfChar(t *testing.T) {
	var testcases = []struct {
		name string
		in   []string
		want [][]string
	}{
		{
			name: "Ok",
			in:   []string{"A", "B", "C"},
			want: [][]string{
				{},
				{"C"},
				{"B"},
				{"C", "B"},
				{"A"},
				{"C", "A"},
				{"B", "A"},
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
			var got = combinationOfChar(tc.in)
			fmt.Println(got)
			if !cmp.Equal(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
