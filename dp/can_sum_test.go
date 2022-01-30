package dp

import "testing"

func TestCanSum(t *testing.T) {
	var testcases = []struct {
		name   string
		target int
		nums   []int
		want   bool
	}{
		{
			name:   "Ok/true",
			target: 7,
			nums:   []int{2, 3, 4, 7},
			want:   true,
		},
		{
			name:   "Ok/fals-2",
			target: 7,
			nums:   []int{20, 13, 4, 2},
		},
		{
			name:   "ok/false-2",
			target: 3001,
			nums:   []int{2, 2, 8, 4, 6},
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got = CanSum(tc.target, tc.nums)
			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
