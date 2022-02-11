package slidingwindows

import "testing"

func TestMinSizeSubarraySum(t *testing.T) {
	var testcases = []struct {
		name string
		nums []int
		sum  int
		want int
	}{
		{
			name: "Ok-1",
			nums: []int{2, 1, 5, 2, 3, 2},
			sum:  7,
			want: 2,
		},
		{
			name: "Ok-2",
			nums: []int{2, 1, 5, 2, 8},
			sum:  7,
			want: 1,
		},
		{
			name: "no-sum-avaiable",
			nums: []int{11, 1, 5, 2, 8},
			sum:  100,
			want: 0,
		},
	}

	for _, tc := range testcases {
		var tc = tc
		t.Run(tc.name, func(t *testing.T) {
			var got = minSizeSubArraySum(tc.nums, tc.sum)

			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
