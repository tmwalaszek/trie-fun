package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	var tt = []struct {
		Key   string
		Value int
	}{
		{"wantt", 2},
		{"want", 3},
		{"abcdert", 90},
		{"abc", 91},
		{"eee", 12},
	}

	trie := NewTrie()

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Insert Key %s with value %d", tc.Key, tc.Value), func(t *testing.T) {
			trie.AddValue(tc.Key, tc.Value)
		})
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Check Key %s value", tc.Key), func(t *testing.T) {
			v := trie.FindKey(tc.Key)
			if tc.Value != v {
				t.Errorf("Wrong value for Key %s. Received value is %d but it should be %d", tc.Key, v, tc.Value)
			}
		})
	}
}
