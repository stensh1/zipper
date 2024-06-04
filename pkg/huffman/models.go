package huffman

import "container/list"

type Ha struct {
	n     int             // number of not equal bytes in input data (often 256) // TODO: refactor to type byte
	freq  [256]int        // array of bytes frequencies
	codes map[byte]string // bytes and their codes made with Huffman tree
	tree  list.List       // list of nodes and then Huffman tree
}

type Node struct {
	ch    byte
	freq  int
	left  *Node
	right *Node
}
