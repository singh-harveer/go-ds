package arrays

import (
	"reflect"
	"testing"
)

func TestPairSum(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name   string
		in     []int
		target int
		want   []int
	}{
		{
			name:   "Ok",
			in:     []int{1, 2, 9, 4},
			target: 10,
			want:   []int{1, 9},
		},

		{
			name:   "not found",
			in:     []int{1, 2, 5, 4},
			target: 8,
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			var got = pairSum(tc.target, tc.in)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
