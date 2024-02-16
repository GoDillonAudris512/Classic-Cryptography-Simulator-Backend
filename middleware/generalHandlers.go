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

func init() {
	for i := uint8('A'); i <= uint8('Z'); i++ {
		alphabetToNumber[i] = int(i - 'A')
		numberToAlphabet[int(i - 'A')] = i
	}
}