package tests

import (
	"fmt"
	"testing"
	"zipper/pkg/lz77"
)

// TestLz77 makes an object of Mtf struct type and sends it to test func
func TestLz77(t *testing.T) {
	o := lz77.Lz77{}

	fmt.Println("LZ77 testing:")
	TestCompressors(t, o)
}
