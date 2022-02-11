package slidingwindows

import "testing"

func TestMaxSumSubArrayOfSizeK(t *testing.T) {
	var testcases = []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Ok-1",
			nums: []int{2, 1, 5, 1, 3, 2},
			k:    3,
			want: 9,
		},
		{
			name: "Ok-2",
			nums: []int{2, 1, 3, 1, 3, 2},
			k:    3,
			want: 7,
		},
		{
			name: "empty",
			nums: []int{},
			k:    3,
			want: 0,
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {

			var got = maxSumSubArrayOfSizeK(tc.k, tc.nums)

			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
