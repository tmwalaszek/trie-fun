package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	var tt = []struct {
		Key     string
		Value   interface{}
		Err     error
		Add     bool
		Check   bool
		Updated bool
	}{

		{"wantt", 2, nil, true, true, false},
		{"want", 3, nil, true, true, false},
		{"abcdert", 90, nil, true, true, false},
		{"abc", 91, nil, true, true, false},
		{"eee", 12, nil, true, true, false},
		{"", 1, ErrEmptyKey, true, true, false},
		{"ee", nil, nil, false, true, false},
		{"e", 11, nil, true, true, false},
		{"eee", 12, nil, false, true, true},
		{"waantt", 34, nil, true, true, false},
		{"ttt", 34, nil, true, false, false},
		{"ttt", 54, nil, true, false, true},
		{"ttt", 54, nil, true, false, true},
	}

	trie := NewTrie()

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Insert Key %s with value %d", tc.Key, tc.Value), func(t *testing.T) {
			if tc.Add {
				updated, err := trie.AddValue(tc.Key, tc.Value)
				if err != tc.Err {
					t.Errorf("Error mismatch: should get %v but got %v", tc.Err, err)
				}

				if updated != tc.Updated {
					t.Errorf("Updated mismatch: should be %v but is %v", tc.Updated, updated)
				}
			}
		})
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Check Key %s value", tc.Key), func(t *testing.T) {
			if !tc.Check {
				return
			}

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
