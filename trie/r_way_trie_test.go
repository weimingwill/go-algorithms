package trie

import (
	"fmt"
	"strings"
)

var shellsST = "she sells sea shells by the sea shore"

func ExampleTrieST() {
	t := NewTrieST()

	parts := strings.Split(shellsST, " ")
	for i, s := range parts {
		t.Put(s, &i)
	}

	fmt.Print(t)

	// Output:
	// 	by 4
	//  sea 6
	//  sells 1
	//  she 0
	//  shells 3
	//  shore 7
	//  the 5
	//  longestPrefixOf("shellsort"):
	//  shells
	//
	//  keysWithPrefix("shor"):
	//  shore
	//
	//  keysThatMatch(".he.l."):
	//  shells
	//
}
