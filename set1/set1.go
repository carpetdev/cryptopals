package set1

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"cryptopals/utils"
)

func Challenge1() {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	answer := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	fmt.Println(utils.Hex2Base64(hex) == answer)
}

func Challenge2() {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	answer := "746865206b696420646f6e277420706c6179"
	fmt.Println(utils.FixedXOR(hex1, hex2) == answer)
}

func Challenge3() {
	hex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	fmt.Println(strings.Join(utils.DecryptXOR(hex), "\n"))
}

func Challenge4() {
	file, err := os.Open("inputs/s1c4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		hex := scanner.Text()
		decryptedStrings := utils.DecryptXOR(hex)
		if len(decryptedStrings) > 0 {
			fmt.Println(hex)
			fmt.Println(strings.Join(decryptedStrings, "\n"))
			fmt.Println()
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
