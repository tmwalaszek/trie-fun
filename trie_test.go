package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	var tt = []struct {
		Key   string
		Value interface{}
		Err   error
		Add   bool
	}{

		{"wantt", 2, nil, true},
		{"want", 3, nil, true},
		{"abcdert", 90, nil, true},
		{"abc", 91, nil, true},
		{"eee", 12, nil, true},
		{"", 1, ErrEmptyKey, true},
		{"ee", nil, nil, false},
	}

	trie := NewTrie()

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Insert Key %s with value %d", tc.Key, tc.Value), func(t *testing.T) {
			if tc.Add {
				err := trie.AddValue(tc.Key, tc.Value)
				if err != tc.Err {
					t.Errorf("Error mismatch: should get %v but got %v", tc.Err, err)
				}
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
				t.Errorf("Wrong value for Key %s Received value is %d but it should be %v", tc.Key, v, tc.Value)
			}
		})
	}
}
