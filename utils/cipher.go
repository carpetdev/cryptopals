package utils

import "math"

// Bruteforces a single-byte XOR cipher
// Uses character frequency to reject non-english text
func DecryptXOR(hexString string) (output []string) {
	if len(hexString)%2 == 1 {
		panic("odd length hex sdsatring")
	}
	n := len(hexString) / 2

	bytes := Hex2Bytes(hexString)
	englishFrequenciesPercent := map[byte]float64{
		'A': 8.2, 'B': 1.5, 'C': 2.8, 'D': 4.3, 'E': 12.7, 'F': 2.2, 'G': 2.0,
		'H': 6.1, 'I': 7.0, 'J': 0.16, 'K': 0.77, 'L': 4.0, 'M': 2.4, 'N': 6.7,
		'O': 7.5, 'P': 1.9, 'Q': 0.12, 'R': 6.0, 'S': 6.3, 'T': 9.1,
		'U': 2.8, 'V': 0.98, 'W': 2.4, 'X': 0.15, 'Y': 2.0, 'Z': 0.074,
	}
KeyLoop:
	for key := 0; key < 256; key++ {
		xorredBytes := make([]byte, n)
		frequencies := make(map[byte]float64)
		for i, b := range bytes {
			xorredByte := b ^ byte(key)
			if xorredByte >= 32 && xorredByte <= 127 { // human-readable
				if xorredByte >= 'a' && xorredByte <= 'z' {
					xorredByte ^= 32 // converts to uppercase
				}
				if xorredByte >= 'A' && xorredByte <= 'Z' {
					frequencies[xorredByte] += 1
				}
			} else {
				continue KeyLoop
			}
			xorredBytes[i] = xorredByte
		}
		chi2 := 0.0
		for b, freq := range frequencies {
			trueFreq := englishFrequenciesPercent[b] * float64(n) / 100
			chi2 += math.Pow(freq-trueFreq, 2) / trueFreq
		}
		if chi2 < 38 { // 0.05 p-value for chi^2_25
			output = append(output, string(xorredBytes[:]))
		}
	}
	return
}
