package tests

import (
	"fmt"
	"testing"
	"zipper/pkg/rle"
)

// TestBwt makes an object of Rle struct type and sends it to test func
func TestRle(t *testing.T) {
	o := rle.Rle{}

	fmt.Println("RLE testing:")
	TestCompressors(t, o)
}
