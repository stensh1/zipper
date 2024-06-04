package huffman

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Decode decodes the input data using a Huffman algorithm and returns decoded data in []byte and error
func (h *Ha) Decode(data []byte) ([]byte, error) {
	h.init()

	if len(data) == 0 {
		return data, nil
	}

	h.readInfo(data) // reading info table
	h.createNodes()  // making the tree's leafs
	h.buildTree()    // making a Huffman tree

	// *5 - 1 byte for char; 2 bytes for frequency and 4 bytes is the first record - number of not equal chars
	return h.readData(data[h.n*5+4:])
}

// readInfo just reads number of not equal bytes and table with byte codes from compressed input
func (h *Ha) readInfo(data []byte) {
	h.n = int(binary.BigEndian.Uint32(data[0:4])) // number of not equal chars from 0 to 3 index

	for i, j := 0, 4; i < h.n; i++ {
		h.freq[data[j]] = int(binary.BigEndian.Uint32(data[j+1 : j+5])) // index j - char, [j+1 : j+5] - char frequency
		j += 5
	}
}

// readData reads bytes from compressed input, reads bits in each one and encodes it in original bytes
func (h *Ha) readData(data []byte) ([]byte, error) {
	var ptr = h.tree.Front().Value.(*Node)
	var result bytes.Buffer // output data
	var cnt = ptr.freq      // how much chars we should decode

	if h.n == 1 {
		for i := 0; i < cnt; i++ {
			if err := result.WriteByte(ptr.ch); err != nil {
				log.Println("Huffman readData decoding one different char failed")
				return nil, err
			}
		}
		return result.Bytes(), nil
	}

	// decoding by the Huffman tree
	for _, b := range data {
		if cnt > 0 {
			for j := 7; j >= 0 && cnt > 0; j-- { // j is bit index positioning
				bit := b & (1 << j) >> j

				if bit == 1 {
					ptr = ptr.right
				} else {
					ptr = ptr.left
				}

				if ptr.left == nil && ptr.right == nil {
					if err := result.WriteByte(ptr.ch); err != nil {
						log.Println("Huffman readData decoding failed")
						return nil, err
					}
					cnt--
					ptr = h.tree.Front().Value.(*Node)
				}
			}
		} else {
			break
		}
	}

	return result.Bytes(), nil
}
