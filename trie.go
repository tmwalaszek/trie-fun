package trie

import (
	"errors"
)

var (
	ErrEmptyKey = errors.New("Empty key")
)

type Trie struct {
	Childrens map[rune]*TrieNode
}

type TrieNode struct {
	RemainingKey string
	Value        interface{}

	Childrens map[rune]*TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Childrens: make(map[rune]*TrieNode),
	}
}

func (t *Trie) FindKey(key string) (interface{}, error) {
	if key == "" {
		return 0, ErrEmptyKey
	}

	return t.findKey(key, nil), nil
}

func (t *Trie) findKey(key string, node *TrieNode) interface{} {
	if node == nil {
		if t.Childrens == nil {
			return nil
		}

		val, ok := t.Childrens[rune(key[0])]
		if !ok {
			return nil
		}

		return t.findKey(key[1:], val)
	}

	if node.RemainingKey == key {
		return node.Value
	}

	if node.Childrens == nil {
		return nil
	}

	val, ok := node.Childrens[rune(key[0])]
	if !ok {
		return nil
	} else {
		return t.findKey(key[1:], val)
	}
}

func (t *Trie) AddValue(key string, value interface{}) error {
	if key == "" {
		return ErrEmptyKey
	}

	t.addValue(key, value, nil)
	return nil
}

func (t *Trie) addValue(key string, value interface{}, node *TrieNode) {
	if node == nil {
		val, ok := t.Childrens[rune(key[0])]
		if !ok {
			t.Childrens[rune(key[0])] = &TrieNode{
				RemainingKey: key[1:],
				Value:        value,
			}
			return
		}

		node = val
		key = key[1:]
	}

	if key == node.RemainingKey {
		node.Value = value
		return
	}

	if key == "" && node.RemainingKey == "" {
		node.Value = value
		return
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
		return
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
		node.Value = 0
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

		return
	}

	t.addValue(key[1:], value, newNode)
}
