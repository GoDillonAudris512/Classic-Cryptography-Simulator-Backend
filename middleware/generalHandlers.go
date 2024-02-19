package middleware

import (
	"fmt"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to Classic Cryptography Simulator Web Service!")
}

// Additional Helper Variables and Functions
var (
	alphabetToNumber = make(map[uint8]int)
	numberToAlphabet = make(map[int]uint8)

	alphabet256ToNumber = make(map[uint8]int)
	numberToAlphabet256 = make(map[int]uint8)

	rotors     = make([][]uint8, 3)
	rotorOrder = make([]int, 3)
	pos    	   = make([]int, 3)
	turnover   = make([]int, 3)
	reflector  = make([]uint8, 26)
	plugboard  = make(map[uint8]uint8)
)

func init() {
	// Dictionary for 26 alphabets
	for i := uint8('A'); i <= uint8('Z'); i++ {
		alphabetToNumber[i] = int(i - 'A')
		numberToAlphabet[int(i-'A')] = i
	}

	// Dictionary for 256 ASCII characters
	for i := uint8(0); i < uint8(255); i++ {
		alphabet256ToNumber[i] = int(i)
		numberToAlphabet256[int(i)] = i
	}
	alphabet256ToNumber[uint8(255)] = 255
	numberToAlphabet256[255] = uint8(255)

	// Variables for Enigma Cipher
	for i := 0; i <= 2; i++ {
		rotors[i] = make([]uint8, 26)
	}
	for i := 0; i < 26; i++ {
		rotors[0][i] = uint8(alphabetToNumber["EKMFLGDQVZNTOWYHXUSPAIBRCJ"[i]])
		rotors[1][i] = uint8(alphabetToNumber["AJDKSIRUXBLHWTMCQGZNPYFVOE"[i]])
		rotors[2][i] = uint8(alphabetToNumber["BDFHJLCPRTXVZNYEIWGAKMUSQO"[i]])
		reflector[i] = uint8(alphabetToNumber["YRUHQSLDPXNGOKMIEBFZCWVJAT"[i]])
	}
	turnover[0] = 16
	turnover[1] = 4
	turnover[2] = 21
}
