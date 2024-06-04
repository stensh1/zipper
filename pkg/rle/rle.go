// Package rle implements a simple lossless RLE compression algorithm
package rle

import "log"

// Encode encodes the input byte slice into a compressed RLE format
// and returns output byte slice
func (r Rle) Encode(data []byte) ([]byte, error) {
	n := len(data)

	// There is no point in compressing
	if n <= 2 {
		return data, nil
	}

	for i := 0; i < n; i++ {
		var cnt byte // count of equal or different bytes next to each other
		// informational byte: 1st bit is flag (1 - next byte repeats n times;
		// 2 - next n bytes are single), 2-8 bits are the value n
		var info byte
		// flag - we are counting equal or different bytes
		var equal bool

		if i < n-1 {
			// if different bytes
			if data[i] != data[i+1] {
				for i < n-1 && data[i] != data[i+1] {
					if cnt >= 128 {
						break
					}

					if err := r.tmp.WriteByte(data[i]); err != nil {
						log.Println("RLE Encode error")
						return nil, err
					}
					cnt++
					i++
					// the condition is not to skip the last byte of data
					if i == n-1 {
						if err := r.tmp.WriteByte(data[i]); err != nil {
							log.Println("RLE Encode error")
							return nil, err
						}
						cnt++
						i++
					}
				}
				i--
				// if equal bytes
			} else {
				cnt = 1
				equal = true
				for i < n-1 && data[i] == data[i+1] {
					// thats 'cause the 1st bit is responsible for
					// the difference between the same bytes and different ones
					// so we can only count less than 128 bytes (2^8 -2^7 - 1)
					if cnt >= 128 {
						break
					}
					cnt++
					i++
				}
			}
		}

		if equal {
			info |= 128
			info |= cnt - 2 // the minimal sequence of identical bytes is 2 (value 0 is 2 bytes)
			if err := r.buf.WriteByte(info); err != nil {
				log.Println("RLE Encode error")
				return nil, err
			}
			if err := r.buf.WriteByte(data[i]); err != nil {
				log.Println("RLE Encode error")
				return nil, err
			}
		} else {
			info |= cnt - 1 // the minimal sequence of different bytes is 1 (value 0 is 1 byte)
			if err := r.buf.WriteByte(info); err != nil {
				log.Println("RLE Encode error")
				return nil, err
			}

			if _, err := r.buf.Write(r.tmp.Bytes()); err != nil {
				log.Println("RLE Encode error")
				return nil, err
			}

			r.tmp.Reset()
		}
	}

	return r.buf.Bytes(), nil
}

// Decode decodes the input RLE byte slice into a decompressed format
// and returns output byte slice
func (r Rle) Decode(data []byte) ([]byte, error) {
	n := len(data)

	// If its empty
	if n == 0 {
		return r.buf.Bytes(), nil
		// If decoding is not required
	} else if n <= 2 && data[0]&128 != 128 {
		return data, nil
	}

	for i := 0; i < n-1; {
		info := data[i]

		// if equal bytes
		if info&128 == 128 {
			for j := 0; j < int((info^128)+2); j++ {
				if err := r.buf.WriteByte(data[i+1]); err != nil {
					log.Println("RLE Decode error")
					return nil, err
				}
			}
			i += 2
			// if different bytes
		} else {
			if _, err := r.buf.Write(data[i+1 : i+1+int(info+1)]); err != nil {
				log.Println("RLE Decode error")
				return nil, err
			}
			i += 1 + int(info+1)
		}
	}

	return r.buf.Bytes(), nil
}
