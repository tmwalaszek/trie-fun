package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	//trie.AddValue("want", 15)
	//trie.AddValue("aaa", 3)
	trie.AddValue("wantt", 2)
	trie.AddValue("want", 3)

	//trie.AddValue("wbl", 10)

	//v := trie.FindKey("want")
	//fmt.Println("Wynik ", v)
	v := trie.FindKey("wantt")
	fmt.Println("Wynik ", v)

	v = trie.FindKey("want")
	fmt.Println("Wynik ", v)

	trie.AddValue("wab", 7)
	v = trie.FindKey("wab")
	fmt.Println("wynik ", v)

	trie.AddValue("wab", 10)
	v = trie.FindKey("wab")
	fmt.Println("wynik ", v)

	trie.AddValue("abcdert", 90)
	v = trie.FindKey("abcdert")
	fmt.Println("wynik ", v)

	trie.AddValue("abcd", 901)
	v = trie.FindKey("abcd")
	fmt.Println("wynik ", v)

	v = trie.FindKey("abcdert")
	fmt.Println("wynik ", v)

}
