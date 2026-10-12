package utils

import (
	"iter"
	"slices"
)

// XORs two equal length hex strings, outputting hex
func FixedXor(hex1, hex2 string) (output string) {
	if len(hex1) != len(hex2) {
		panic("hex strings are different lengths")
	}

	bytes := slices.Collect(Values(Hex2Bytes(hex1)))
	for i, b := range Hex2Bytes(hex2) {
		bytes[i] ^= b
	}

	alphabet := "0123456789abcdef"
	for _, b := range bytes {
		output += string(alphabet[b>>4]) + string(alphabet[b&15])
	}
	return
}

func SingleByteXor(seq iter.Seq2[int, byte], b byte) iter.Seq2[int, byte] {
	return func(yield func(int, byte) bool) {
		for i, v := range seq {
			if !yield(i, v^b) {
				return
			}
		}
	}
}
