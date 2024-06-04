// Package huffman implements compression and decompression algorithms using the Huffman method.
package huffman

import (
	"container/list"
)

// Init just refresh Ha struct fields 'cause of all methods *Ha didn't work in goroutines w/ refreshing
func (h *Ha) init() {
	h.n = 0
	h.freq = [256]int{}
	h.tree = list.List{}
	h.codes = make(map[byte]string)
}

// createNodes makes sorted list of bytes from data with their frequents
func (h *Ha) createNodes() {
	for i, ch := range h.freq {
		if ch != 0 {
			node := Node{byte(i), h.freq[i], nil, nil}
			h.insertNode(&node)
		}
	}
	h.n = h.tree.Len() // number of created nodes
}

// buildTree is making Huffman binary tree
func (h *Ha) buildTree() {
	n := h.n // number of not equal chars

	// while not a one root
	for n > 1 {
		left := h.tree.Back()
		h.tree.Remove(left)
		right := h.tree.Back()
		h.tree.Remove(right)

		node := Node{0, left.Value.(*Node).freq + right.Value.(*Node).freq, left.Value.(*Node),
			right.Value.(*Node)}
		right, left = nil, nil

		h.insertNode(&node)
		n -= 1
	}
}

// insertNode is making sorted insertion into list of nodes
func (h *Ha) insertNode(n *Node) {
	var ptr *list.Element

	// if new tree
	if h.tree.Len() == 0 {
		h.tree.PushFront(n)
		return
	}

	// searching the right place. Sorting from small to big.
	for ptr = h.tree.Front(); ptr != h.tree.Back() && ptr.Value.(*Node).freq >= n.freq; ptr = ptr.Next() {
	}

	h.tree.InsertBefore(n, ptr)
}
