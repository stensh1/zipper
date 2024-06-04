package tests

import (
	"testing"
)

// Zipper is an interface for all compression/decompression algs
type Zipper interface {
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

// TestCompressors is a main func for all de/compression algs that takes object of type Zipper as argument
func TestCompressors(t *testing.T, z Zipper) {
	//The set of testing data
	tests := []struct {
		name         string
		in, expected []byte
	}{
		{"Empty file", []byte(""), []byte("")},
		{"Single byte", []byte("0"), []byte("0")},
		{"Two different bytes", []byte("ab"), []byte("ab")},
		{"Two equal bytes", []byte("aa"), []byte("aa")},
		{"Three different bytes", []byte("abc"), []byte("abc")},
		{"Three equal bytes", []byte("aaa"), []byte("aaa")},
		{"Random string", []byte("aaaaaaaaaaaaaaaaaaaaaaabbbbbdweewfewqqqqqqqqqqqqqqqqqqqq111234gff"),
			[]byte("aaaaaaaaaaaaaaaaaaaaaaabbbbbdweewfewqqqqqqqqqqqqqqqqqqqq111234gff")},
	}

	// Running goroutine for each dataset
	for _, i := range tests {
		t.Run(i.name, func(t *testing.T) {
			zip := z
			encoded, err := zip.Encode(i.in)
			if err != nil {
				t.Errorf("Should not produce an error")
			}

			decoded, err := zip.Decode(encoded)
			if err != nil {
				t.Errorf("Should not produce an error")
			}

			if len(decoded) == len(i.expected) {
				for j := range decoded {
					if i.expected[j] != decoded[j] {
						t.Errorf("Result was incorrect, got: %s, want: %s.", decoded, i.expected)
					}
				}
			} else {
				t.Errorf("Result was an incorrect, got: %s, want: %s.", decoded, i.expected)
			}
		})
	}
}
