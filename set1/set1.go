package set1

import (
	"cryptopals/bitwise"
	"cryptopals/encoding"
	"fmt"
)

func Challenge1() {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	fmt.Println(encoding.Hex2Base64(hex) == answer)
}

func Challenge2() {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	answer := "746865206b696420646f6e277420706c6179"
	fmt.Println(bitwise.FixedXOR(hex1, hex2) == answer)
}
}
