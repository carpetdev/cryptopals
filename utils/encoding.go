package utils

import "iter"

func Hex2Bytes(hexString string) iter.Seq2[int, byte] {
	if len(hexString)%2 == 1 {
		panic("odd length hex string")
	}

	hex2Byte := map[byte]byte{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'a': 10, 'A': 10, 'b': 11, 'B': 11, 'c': 12,
		'C': 12, 'd': 13, 'D': 13, 'e': 14, 'E': 14, 'f': 15, 'F': 15,
	} // see https://cs.opensource.google/go/go/+/refs/tags/go1.27.2:src/encoding/hex/hex.go;l=18

	return func(yield func(int, byte) bool) {
		for i := 0; i < len(hexString)/2; i++ {
			if !yield(i, hex2Byte[hexString[2*i]]<<4|hex2Byte[hexString[2*i+1]]) { // TODO: handle incorrect hex (currently picks 0 i think)
				return
			}
		}
	}
}

func Hex2Base64(hexString string) (base64String string) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	var prevSegment, offset byte
	for _, hexByte := range Hex2Bytes(hexString) {
		newSegment := hexByte >> (2 + offset)
		sextet := (prevSegment << (6 - offset)) | newSegment
		// fmt.Println(offset)
		// fmt.Printf("%b\n", [...]byte{hexByte, prevSegment, newSegment, sextet})
		base64String += string(alphabet[sextet])
		prevSegment = hexByte ^ (newSegment << (2 + offset))
		offset += 2
		if offset == 6 {
			base64String += string(alphabet[prevSegment])
			offset = 0
			prevSegment = 0
		}
	}
	if offset != 0 {
		base64String += string(alphabet[prevSegment<<(6-offset)])
		switch offset {
		case 2:
			base64String += "=="
		case 4:
			base64String += "="
		default:
			panic("invalid offset")
		}
	}
	return
}
