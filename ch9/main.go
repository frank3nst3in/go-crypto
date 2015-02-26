package main

import (
	"bytes"
	"fmt"
)

func pad_block(block []byte, pad int) []byte {
	size := len(block)
	switch {
	case pad == size:
		return block
	case pad < size:
		return block[:pad]
	default:
		n := pad - size%pad
		padded := bytes.Repeat([]byte{byte(n)}, n)
		return append(block, padded...)
	}
}

func main() {
	block := []byte("YELLOW SUBMARINE")
	pad := 20
	fmt.Printf("%q\n", pad_block(block, pad))
}
