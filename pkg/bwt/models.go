package bwt

import "bytes"

type Bwt struct {
	result bytes.Buffer // data output slice
	n      int          // input data length
}

type Pair struct {
	char  byte
	index int
}
