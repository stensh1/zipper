package tests

import (
	"fmt"
	"testing"
	"zipper/pkg/bwt"
)

// TestBwt makes an object of Bwt struct type and sends it to test func
func TestBwt(t *testing.T) {
	o := bwt.Bwt{}

	fmt.Println("BWT testing:")
	TestCompressors(t, o)
}
