package bitwise

import (
	"cryptopals/encoding"
	"slices"
)

// XORs two equal length hex strings, outputting hex
func FixedXOR(hex1, hex2 string) string {
	if len(hex1) != len(hex2) {
		panic("hex strings are different lengths")
	}

	outputBytes := slices.Collect(encoding.Hex2Bytes(hex1))
	i := 0
	for b := range encoding.Hex2Bytes(hex2) {
		outputBytes[i] ^= b
		i++
	}

	alphabet := "0123456789abcdef"
	output := ""
	for _, b := range outputBytes {
		output += string(alphabet[b>>4]) + string(alphabet[b&15])
	}
	return output
}
