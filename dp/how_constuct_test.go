package dp

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHowConstruct(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name   string
		target string
		in     []string
		want   []string
	}{
		{
			name:   "Ok/true",
			target: "abcdef",
			in:     []string{"ab", "a", "bc", "abc", "def"},
			want:   []string{"abc", "def"},
		},

		{
			name:   "Ok/false",
			target: "abcdefg",
			in:     []string{"ab", "a", "bc", "abc", "def"},
			want:   []string{},
		},

		{
			name:   "test-timeout-without-memoization",
			target: "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
			in:     []string{"e", "eeeee", "eeee", "e", "eeee"},
			want:   []string{},
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var got = HowConstruct(tc.target, tc.in)
			if cmp.Equal(got, tc.want) {
				t.Fatalf("want %v got %v", tc.want, got)
			}
		})
	}
}
