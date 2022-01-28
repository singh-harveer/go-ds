package dp

import "testing"

func TestCanConstruct(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name   string
		target string
		in     []string
		want   bool
	}{
		{
			name:   "Ok/true",
			target: "abcdef",
			in:     []string{"ab", "a", "bc", "abc", "def"},
			want:   true,
		},

		{
			name:   "Ok/false",
			target: "abcdefg",
			in:     []string{"ab", "a", "bc", "abc", "def"},
			want:   false,
		},

		{
			name:   "test-timeout-without-memoization",
			target: "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
			in:     []string{"e", "eeeee", "eeee", "e", "eeee"},
			want:   false,
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got = CanConstruct(tc.target, tc.in)

			if got != tc.want {
				t.Fatalf("want %v got %v", tc.want, got)
			}

		})
	}
}
