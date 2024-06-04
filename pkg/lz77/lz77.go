// Package lz77 implements compression and decompression lz77 algorithms.
package lz77

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Encode implements compression of []byte using the LZ77 algorithm and returns []byte and error
func (lz Lz77) Encode(data []byte) ([]byte, error) {
	var encoded []Lz77
	var result bytes.Buffer // output data
	var tmp [4]byte         // buffer for converting int type to [4]byte
	var windowSize = 32

	n := len(data)

	for i := 0; i < n; {
		var matchOffset, matchLength int
		for j := 1; j <= windowSize && i-j >= 0; j++ {
			length := 0
			for length < windowSize && i+length < n && data[i-j+length] == data[i+length] {
				length++
			}
			if length > matchLength {
				matchOffset = j
				matchLength = length
			}
		}

		var nextChar byte
		if i+matchLength < n {
			nextChar = data[i+matchLength]
		}
		encoded = append(encoded, Lz77{matchOffset, matchLength, nextChar})
		i += matchLength + 1
	}

	// converting pairs slice into byte slice
	for _, k := range encoded {
		binary.BigEndian.PutUint32(tmp[:], uint32(k.Offset))
		if _, err := result.Write(tmp[:]); err != nil {
			log.Println("LZ77 Encode offset converting error")
			return nil, err
		}
		binary.BigEndian.PutUint32(tmp[:], uint32(k.Length))
		if _, err := result.Write(tmp[:]); err != nil {
			log.Println("LZ77 Encode length converting error")
			return nil, err
		}

		if err := result.WriteByte(k.Next); err != nil {
			log.Println("LZ77 Encode writing next char error")
			return nil, err
		}
	}

	return result.Bytes(), nil
}

// Decode implements decoding of a compressed []byte using the LZ77 algorithm and returns decompressed []byte and error
func (lz Lz77) Decode(data []byte) ([]byte, error) {
	var result bytes.Buffer // output data

	for i := 0; i < len(data); {
		var offset = int(binary.BigEndian.Uint32(data[i : i+4]))
		var length = int(binary.BigEndian.Uint32(data[i+4 : i+8]))
		var next = data[i+8]
		i += 9

		if offset == 0 && length == 0 {
			if err := result.WriteByte(next); err != nil {
				log.Println("LZ77 Decode writing next char error")
				return nil, err
			}
		} else {
			start := result.Len() - offset
			for i := 0; i < length; i++ {
				if err := result.WriteByte(result.String()[start+i]); err != nil {
					log.Println("LZ77 Decode writing char error")
					return nil, err
				}
			}
			if next != 0 {
				if err := result.WriteByte(next); err != nil {
					log.Println("LZ77 Decode writing next char error")
					return nil, err
				}
			}
		}
	}

	return result.Bytes(), nil
}
