package trie

import (
	"errors"
	"sync"
)

var (
	ErrEmptyKey = errors.New("Empty key")
)

type Trie struct {
	Childrens map[rune]*TrieNode

	mutex *sync.RWMutex
}

type TrieNode struct {
	RemainingKey string
	Value        interface{}

	Childrens map[rune]*TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Childrens: make(map[rune]*TrieNode),
		mutex:     &sync.RWMutex{},
	}
}

func (t *Trie) FindKey(key string) (interface{}, error) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if key == "" {
		return 0, ErrEmptyKey
	}

	return t.findKey(key), nil
}

type WalkFn func() (*TrieNode, interface{})

func (t *Trie) walk(key string) WalkFn {
	var node *TrieNode
	var ok bool

	f := func() (*TrieNode, interface{}) {
		if node == nil {
			if t.Childrens == nil {
				return nil, nil
			}

			node, ok = t.Childrens[rune(key[0])]
			if !ok {
				return nil, nil
			}

			key = key[1:]
			return node, nil
		}

		if node.RemainingKey == key {
			return nil, node.Value
		}

		if node.Childrens == nil {
			return nil, nil
		}

		node, ok = node.Childrens[rune(key[0])]
		if !ok {
			return nil, nil
		}
		key = key[1:]

		return node, nil
	}

	return f
}

func (t *Trie) findKey(key string) interface{} {
	walkFn := t.walk(key)

	for {
		node, value := walkFn()
		if node == nil && value == nil {
			return nil
		} else if value != nil {
			return value
		}
	}
}

func (t *Trie) AddValue(key string, value interface{}) (bool, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if key == "" {
		return false, ErrEmptyKey
	}

	return t.addValue(key, value), nil
}

func (t *Trie) addValue(key string, value interface{}) bool {
	var node *TrieNode
	var ok bool

	for {
		if node == nil {
			node, ok = t.Childrens[rune(key[0])]
			if !ok {
				t.Childrens[rune(key[0])] = &TrieNode{
					RemainingKey: key[1:],
					Value:        value,
				}

				return false
			}

			key = key[1:]
		}

		var updated bool
		if key == node.RemainingKey {
			if node.Value != nil {
				updated = true
			}

			node.Value = value
			return updated
		}

		if key == "" && node.RemainingKey == "" {
			if node.Value != nil {
				updated = true
			}

			node.Value = value
			return updated
		}

		if key == "" && node.RemainingKey != "" {
			if node.Childrens == nil {
				node.Childrens = make(map[rune]*TrieNode)
			}

			node.Childrens[rune(node.RemainingKey[0])] = &TrieNode{
				RemainingKey: node.RemainingKey[1:],
				Value:        node.Value,
			}

			node.RemainingKey = ""
			node.Value = value
			return updated
		}

		var newNode *TrieNode
		if node.Childrens == nil {
			node.Childrens = make(map[rune]*TrieNode)

			node.Childrens[rune(node.RemainingKey[0])] = &TrieNode{
				RemainingKey: node.RemainingKey[1:],
				Value:        node.Value,
			}

			newNode = node.Childrens[rune(node.RemainingKey[0])]

			node.RemainingKey = ""
			node.Value = nil
		}

		if newNode == nil {
			newNode = node.Childrens[rune(key[0])]
		}

		_, ok := node.Childrens[rune(key[0])]
		if !ok {
			node.Childrens[rune(key[0])] = &TrieNode{
				RemainingKey: key[1:],
				Value:        value,
			}

			return false
		}

		key = key[1:]
		node = newNode
	}
}
