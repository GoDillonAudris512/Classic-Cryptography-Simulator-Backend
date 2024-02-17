package middleware

import (
	"encoding/json"
	"net/http"

	"classic-crypt/model"
)

func ExtendedEuclidean(a int, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}

	gcd, x1, y1 := ExtendedEuclidean(b, a%b)
	x := y1
	y := x1 - (a / b * y1)

	return gcd, x, y
}

func HandleAffine(response http.ResponseWriter, request *http.Request) {
	var reqToken model.AffineRequestToken
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqToken)

	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	gcd, x, _ := ExtendedEuclidean(reqToken.Slope, 26)

	if gcd != 1 {
		response.Header().Set("Content-Type", "application/json")

		var resToken model.AffineResponseToken
		resToken.Success = false
		resToken.Output = "Unable to process because the value of slope is not coprime to the size of alphabet (26)"

		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(resToken)
		return
	}

	if reqToken.Encrypt {
		EncryptAffine(reqToken.Input, reqToken.Slope, reqToken.Intercept, gcd, x, response)
	} else {
		DecryptAffine(reqToken.Input, reqToken.Slope, reqToken.Intercept, gcd, x, response)
	}
}

func EncryptAffine(input string, slope int, intercept int, gcd int, x int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	cipherText := ""
	for _, char := range input {
		token := alphabetToNumber[uint8(char)]

		cipherToken := numberToAlphabet[((slope*token)+intercept)%26]
		cipherText += string(cipherToken)
	}

	var resToken model.AffineResponseToken
	resToken.Success = true
	resToken.Output = cipherText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func DecryptAffine(input string, slope int, intercept int, gcd int, x int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	slopeInverse := (x + 26) % 26

	plainText := ""
	for _, char := range input {
		token := alphabetToNumber[uint8(char)]

		result := token - intercept
		for result < 0 {
			result += 26
		}
		
		cipherToken := numberToAlphabet[(slopeInverse*(result))%26]
		plainText += string(cipherToken)
	}

	var resToken model.AffineResponseToken
	resToken.Success = true
	resToken.Output = plainText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}
