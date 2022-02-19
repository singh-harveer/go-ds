package bitmanipulations

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsEven(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "Ok",
			num:  2,
			want: true,
		},
		{
			name: "Ok",
			num:  3,
			want: false,
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			var got = isEven(tc.num)
			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name string
		num1 int
		num2 int
		want []int
	}{
		{
			name: "Ok",
			num1: 2,
			num2: 3,
			want: []int{2, 3},
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			var got = swap(tc.num1, tc.num2)
			if cmp.Equal(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestGetKthBit(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name string
		num  int
		k    int
		want int
	}{
		{
			name: "Ok-1",
			num:  5,
			k:    1,
			want: 0,
		},
		{
			name: "Ok-2",
			num:  8,
			k:    3,
			want: 1,
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			var got = getKthBits(tc.num, tc.k)
			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestBinaryToDecimal(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		name string
		num  []int
		want int
	}{
		{
			name: "Ok-1",
			num:  []int{1, 0, 1},
			want: 5,
		},
		{
			name: "Ok-1",
			num:  []int{1, 0, 0, 0},
			want: 8,
		},
		{
			name: "Ok-1",
			num:  []int{0, 0, 0, 0},
			want: 0,
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			var got = binaryToDecimal(tc.num)
			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
