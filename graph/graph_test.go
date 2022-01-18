package graph_test

import (
	"golang/ds/graph"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDepthFirstTraversal(t *testing.T) {
	t.Parallel()
	var testcases = []struct {
		name   string
		graph  map[int][]int
		source int
		want   []int
	}{
		{
			name: "Ok-1",
			graph: map[int][]int{
				1: []int{2, 3},
				2: []int{4},
				3: []int{5},
				5: []int{4},
			},
			source: 1,
			want:   []int{1, 3, 5, 4, 2, 4},
		},
		{
			name: "Ok-2",
			graph: map[int][]int{
				1: []int{2, 3},
				2: []int{4},
				3: []int{5},
				4: []int{6},
			},
			source: 2,
			want:   []int{2, 4, 6},
		},
	}

	for _, tc := range testcases {
		var tc = tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			var got = graph.DepthFirstTraversal(tc.source, tc.graph)
			if !cmp.Equal(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestBreadthFirstTraversal(t *testing.T) {
	t.Parallel()
	var testcases = []struct {
		name   string
		graph  map[int][]int
		source int
		want   []int
	}{
		{
			name: "Ok-1",
			graph: map[int][]int{
				1: []int{2, 3},
				2: []int{4},
				3: []int{5},
				5: []int{4},
			},
			source: 1,
			want:   []int{1, 2, 3, 4, 5, 4},
		},
		{
			name: "Ok-2",
			graph: map[int][]int{
				1: []int{2, 3},
				2: []int{4},
				3: []int{5},
				4: []int{6},
			},
			source: 2,
			want:   []int{2, 4, 6},
		},
	}

	for _, tc := range testcases {
		var tc = tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got = graph.BreathFirstTraversal(tc.source, tc.graph)
			if !cmp.Equal(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestHasPath(t *testing.T) {
	t.Parallel()
	var testcases = []struct {
		name   string
		graph  map[int][]int
		source int
		dest   int
		want   bool
	}{
		{
			name: "Ok-1",
			graph: map[int][]int{
				1: []int{2, 3},
				2: []int{4},
				3: []int{5},
				5: []int{4},
			},
			source: 1,
			dest:   5,
			want:   true,
		},
		{
			name: "Ok-2",
			graph: map[int][]int{
				1: []int{2, 3},
				2: []int{4},
				3: []int{5},
				4: []int{6},
			},
			source: 2,
			dest:   5,
			want:   false,
		},
	}

	for _, tc := range testcases {
		var tc = tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var got = graph.HasPath(tc.source, tc.dest, tc.graph)
			if got != tc.want {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
