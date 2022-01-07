package trie_test

import (
	"golang/ds/trie"
	"testing"
)

func TestTrieInsertAndSearch(t *testing.T) {
	t.Parallel()
	//  create new trie object.
	var trieObj = trie.New()

	var testcases = []struct {
		name       string
		in         string
		searchWord string
		want       bool
	}{
		{
			name:       "Ok",
			in:         "hello",
			searchWord: "hello",
			want:       true,
		},
		{
			name:       "notFound",
			in:         "hello",
			searchWord: "not hello",
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			trieObj.Insert(tc.in)
			var got = trieObj.Search(tc.searchWord)

			if tc.want != got {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
