package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	var tt = []struct {
		Key   string
		Value int
		Err   error
	}{
		{"wantt", 2, nil},
		{"want", 3, nil},
		{"abcdert", 90, nil},
		{"abc", 91, nil},
		{"eee", 12, nil},
		{"", 1, ErrEmptyKey},
	}

	trie := NewTrie()

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Insert Key %s with value %d", tc.Key, tc.Value), func(t *testing.T) {
			err := trie.AddValue(tc.Key, tc.Value)
			if err != tc.Err {
				t.Errorf("Error mismatch: should get %v but got %v", tc.Err, err)
			}
		})
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Check Key %s value", tc.Key), func(t *testing.T) {
			v, err := trie.FindKey(tc.Key)
			if err != tc.Err {
				t.Fatalf("Error mismatch: should get %v but got %v", tc.Err, err)
			}

			if err != nil {
				return
			}

			if tc.Value != v {
				t.Errorf("Wrong value for Key %s Received value is %d but it should be %d", tc.Key, v, tc.Value)
			}
		})
	}
}
