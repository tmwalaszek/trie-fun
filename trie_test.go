package trie

import (
	"fmt"
	"testing"
	"time"

	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func createTree(keys, length int) {
	t := NewTrie()

	for i := 0; i < keys; i++ {
		t.AddValue(StringWithCharset(length, charset), i)
	}
}

func benchmarkCreateTree(keys, length int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		createTree(keys, length)
	}
}

func BenchmarkCreateTree_10_10(b *testing.B) {
	benchmarkCreateTree(10, 10, b)
}

func BenchmarkCreateTree_10_100(b *testing.B) {
	benchmarkCreateTree(10, 100, b)
}

func BenchmarkCreateTree_10_1000(b *testing.B) {
	benchmarkCreateTree(10, 1000, b)
}

func BenchmarkCreateTree_100_10(b *testing.B) {
	benchmarkCreateTree(100, 10, b)
}

func BenchmarkCreateTree_100_100(b *testing.B) {
	benchmarkCreateTree(100, 100, b)
}

func BenchmarkCreateTree_100_1000(b *testing.B) {
	benchmarkCreateTree(100, 1000, b)
}

func BenchmarkCreateTree_1000_10(b *testing.B) {
	benchmarkCreateTree(100, 10, b)
}

func BenchmarkCreateTree_1000_100(b *testing.B) {
	benchmarkCreateTree(1000, 100, b)
}

func BenchmarkCreateTree_1000_1000(b *testing.B) {
	benchmarkCreateTree(10000, 1000, b)
}

func BenchmarkCreateTree_10000_10(b *testing.B) {
	benchmarkCreateTree(10000, 10, b)
}

func BenchmarkCreateTree_10000_100(b *testing.B) {
	benchmarkCreateTree(10000, 100, b)
}

func BenchmarkCreateTree_10000_1000(b *testing.B) {
	benchmarkCreateTree(10000, 1000, b)
}

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

	for _, tc := range tt {
		if tc.Add {
			err := trie.DeleteKey(tc.Key)
			if err != nil && tc.Key != "" {
				t.Errorf("Error removing key %s: %v", tc.Key, err)
			}
		}
	}
}
