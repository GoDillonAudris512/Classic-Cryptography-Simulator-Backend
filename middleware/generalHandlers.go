package middleware

import (
	"fmt"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, "Hello, welcome to Classic Cryptography Simulator Web Service!")
}

// Additional Helper Variables and Functions
var alphabetToNumber = make(map[uint8]int)
var numberToAlphabet = make(map[int]uint8)
var alphabet256ToNumber = make(map[uint8]int)
var numberToAlphabet256 = make(map[int]uint8)

func init() {
	for i := uint8('A'); i <= uint8('Z'); i++ {
		alphabetToNumber[i] = int(i - 'A')
		numberToAlphabet[int(i - 'A')] = i
	}

	for i := uint8(0); i < uint8(255); i++ {
		alphabet256ToNumber[i] = int(i)
		numberToAlphabet256[int(i)] = i
	}
	alphabet256ToNumber[uint8(255)] = 255
	numberToAlphabet256[255] = uint8(255)

}