package huffman

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Encode converts the input data using a Huffman algorithm and returns encoded data in []byte and error
func (h *Ha) Encode(data []byte) ([]byte, error) {
	var result bytes.Buffer

	h.init()

	if len(data) == 0 {
		return data, nil
	}

	h.countFreq(data)                              // builds bytes frequents from input data
	h.createNodes()                                // making the tree's leafs
	h.buildTree()                                  // making a Huffman tree
	h.buildCodes(h.tree.Front().Value.(*Node), "") // encoding bytes

	// Writing info table
	tmp, err := h.writeInfo()
	if err != nil {
		log.Println("Huffman Encode writing table error")
		return nil, err
	}
	if _, err := result.Write(tmp); err != nil {
		log.Println("Huffman Encode writing table error")
		return nil, err
	}

	// Writing encoded data
	tmp, err = h.writeData(data)
	if err != nil {
		log.Println("Huffman Encode writing data error")
		return nil, err
	}
	if _, err := result.Write(tmp); err != nil {
		log.Println("Huffman Encode writing data error")
		return nil, err
	}

	return result.Bytes(), nil
}

// countFreq builds bytes frequents from input data
func (h *Ha) countFreq(data []byte) {
	for _, ch := range data {
		h.freq[ch]++
	}
}

// buildCodes is alg of inorder bypass
func (h *Ha) buildCodes(ptr *Node, code string) {
	// case if we have only one different char, just set it to zero
	if h.n == 1 {
		h.codes[ptr.ch] = "0"
		return
	}

	// last child found
	if ptr.left == nil && ptr.right == nil {
		h.codes[ptr.ch] = code
	} else { // recursive bypass
		h.buildCodes(ptr.left, code+"0")
		h.buildCodes(ptr.right, code+"1")
	}
}

// writeInfo is writing number of not equal bytes in input data and table of byte codes
func (h *Ha) writeInfo() ([]byte, error) {
	var result bytes.Buffer // output data
	var b [4]byte           // buffer to translate int to 4 bytes

	binary.BigEndian.PutUint32(b[:], uint32(h.n)) // (int) number of not equal bytes to [4]byte
	if _, err := result.Write(b[:]); err != nil {
		log.Println("Huffman writeInfo writing number of nodes error")
		return nil, err
	}

	// writing frequency table
	for i, k := range h.freq {
		if k != 0 {
			if err := result.WriteByte(byte(i)); err != nil {
				log.Println("Huffman writeInfo writing frequency table char error")
				return nil, err
			}
			binary.BigEndian.PutUint32(b[:], uint32(k))
			if _, err := result.Write(b[:]); err != nil {
				log.Println("Huffman writeInfo writing frequency table frequency error")
				return nil, err
			}
		}
	}

	return result.Bytes(), nil
}

// writeData is writing codes for one by one byte from input
func (h *Ha) writeData(data []byte) ([]byte, error) {
	var bits bytes.Buffer
	var result []byte

	for _, ch := range data {
		if _, err := bits.Write([]byte(h.codes[ch])); err != nil {
			log.Println("Huffman writeData writing char failed")
			return nil, err
		}
	}

	result, err := writeBits(bits.Bytes())
	if err != nil {
		log.Println("Huffman writeData writing code failed")
		return nil, err
	}

	return result, nil
}

// writeBits accepts a byte slice consisting of only 1 and 0 and writes them as bits of the byte output []byte slice
func writeBits(data []byte) ([]byte, error) {
	var result bytes.Buffer // output data
	var currentByte byte    // building byte
	var bitIndex int        // bit position in currentByte

	for _, bit := range data {
		// if bit is 1
		if bit == 49 {
			currentByte |= 1 << (7 - bitIndex)
		}

		bitIndex++

		// new byte if current is full
		if bitIndex == 8 {
			if err := result.WriteByte(currentByte); err != nil {
				log.Println("Huffman writeBits error")
				return nil, err
			}
			currentByte = 0
			bitIndex = 0
		}
	}

	// end of data but we have not fulfilled byte
	if bitIndex > 0 {
		if err := result.WriteByte(currentByte); err != nil {
			log.Println("Huffman writeBits error")
			return nil, err
		}
	}

	return result.Bytes(), nil
}
