package tests

import (
	"fmt"
	"testing"
	"zipper/pkg/mtf"
)

// TestMtf makes an object of Mtf struct type and sends it to test func
func TestMtf(t *testing.T) {
	o := mtf.Mft{}

	fmt.Println("MTF testing:")
	TestCompressors(t, o)
}
