// Package bwt implements the BVT algorithm and its inverse for reordering input data for further compression
package bwt

import (
	"log"
	"sort"
)

// Encode converts the input data using an algorithm BWT and returns encoded data in []byte and error
func (b Bwt) Encode(data []byte) ([]byte, error) {
	input := string(data)
	input += "$" // don't have to save original table index
	b.n = len(input)

	// Generating all of cyclic line shifts
	table := make([]string, b.n)
	for i := range table {
		table[i] = input[i:] + input[:i]
	}

	// Sorting lexicographically
	sort.Strings(table)

	// Making last column string
	for _, row := range table {
		if err := b.result.WriteByte(row[b.n-1]); err != nil {
			log.Println("BWT Encode error")
			return nil, err
		}
	}

	return b.result.Bytes(), nil
}

// Decode decodes the input data using an algorithm BWT and returns decoded data in []byte and error
func (b Bwt) Decode(data []byte) ([]byte, error) {
	var index int
	b.n = len(data)

	// Making pairs array
	table := make([]Pair, b.n)
	for i := range data {
		table[i] = Pair{data[i], i}
	}

	// Sorting lexicographically and then by the index
	sort.SliceStable(table, func(i, j int) bool {
		return table[i].char < table[j].char
	})

	// To find start point
	for _, k := range table {
		if k.char == '$' {
			index = k.index
			break
		}
	}

	// Decoding
	for i := 0; i < b.n-1; i++ {
		if err := b.result.WriteByte(table[index].char); err != nil {
			log.Println("BWT Decode error")
			return nil, err
		}
		index = table[index].index
	}

	return b.result.Bytes(), nil
}
