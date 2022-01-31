package dp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHowSum(t *testing.T) {
	var testcases = []struct {
		name   string
		target int
		nums   []int
		want   []int
	}{
		{
			name:   "Ok",
			target: 7,
			nums:   []int{2, 3, 4, 7},
			want:   []int{2, 2, 3},
		},
		{
			name:   "sum-not-pssible-1",
			target: 7,
			nums:   []int{},
		},
		{
			name:   "sum-not-possible-2",
			target: 3001,
			nums:   []int{2, 2, 8, 4, 6},
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			var got = HowSum(tc.target, tc.nums)

			if !cmp.Equal(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
