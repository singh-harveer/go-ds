package slidingwindows

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTripletsSumToZero(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name string
		in   []int
		want [][]int
	}{
		{
			name: "Ok",
			in:   []int{-3, 0, 1, 2, -1, 1, -2},
			want: [][]int{{-3, 1, 2}, {-2, 0, 2}, {-2, 1, 1}, {-1, 0, 1}},
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			var got = AllTriplets(tc.in)

			if !cmp.Equal(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
