package trie

// TST represents ternary search trie.
type TST struct {
	root TNode
}

type TNode struct {
	value *int // it can also be an object
	c     byte
	left  *TNode
	mid   *TNode
	right *TNode
}

func (t *TST) Get(s string) *int {

}

func (t *TST) Put(s string, val *int) {

}

func (t *TST) Contains(s string) bool {
}
