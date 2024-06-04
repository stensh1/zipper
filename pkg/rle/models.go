package rle

import "bytes"

type Rle struct {
	buf bytes.Buffer
	tmp bytes.Buffer
}
