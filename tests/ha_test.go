package tests

import (
	"fmt"
	"testing"
	"zipper/pkg/huffman"
)

// TestHa makes an object of Ha struct type and sends it to test func
func TestHa(t *testing.T) {
	o := &huffman.Ha{}

	fmt.Println("Huffman testing:")
	TestCompressors(t, o)
}
