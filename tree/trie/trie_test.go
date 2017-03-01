package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	paths := []string{"aaaa", "aaab", "bbbb", "ccca", "aab", "bba", "abc"}
	for _, path := range paths {
		trie.Insert(path)
		if !trie.Has(path) {
			t.Fail()
		}
	}

	for _, path := range paths {
		trie.Delete(path)
		if trie.Has(path) {
			t.Fail()
		}
	}

	if len(trie.Root.Children) > 0 {
		t.Fail()
	}
}
