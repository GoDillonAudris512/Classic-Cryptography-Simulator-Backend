package middleware

import (
	"encoding/json"
	"net/http"

	"classic-crypt/model"
)

func BuildKeyText(text string, key string) string {
	textLength := len(text)
	keyLength := len(key)

	if keyLength >= textLength {
		return key[:textLength]
	} else {
		keyText := key

		for len(keyText) < textLength {
			keyText += key
		}

		return keyText[:textLength]
	}
}

func HandleVigenere(response http.ResponseWriter, request *http.Request) {
	var reqToken model.VigenereRequestToken
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqToken)

	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	keyText := BuildKeyText(reqToken.Input, reqToken.Key)

	if reqToken.Encrypt {
		EncryptVigenere(reqToken.Input, keyText, response)
	} else {
		DecryptVigenere(reqToken.Input, keyText, response)
	}
}

func EncryptVigenere(input string, keyText string, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	cipherText := ""
	for i := 0; i < len(input); i++ {
		token1 := alphabetToNumber[input[i]]
		token2 := alphabetToNumber[keyText[i]]

		cipherToken := numberToAlphabet[(token1+token2)%26]
		cipherText += string(cipherToken)
	}

	var resToken model.VigenereResponseToken
	resToken.Success = true
	resToken.Output = cipherText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func DecryptVigenere(input string, keyText string, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	plainText := ""
	for i := 0; i < len(input); i++ {
		token1 := alphabetToNumber[input[i]]
		token2 := alphabetToNumber[keyText[i]]

		plainToken := numberToAlphabet[(token1-token2+26)%26]
		plainText += string(plainToken)
	}

	var resToken model.VigenereResponseToken
	resToken.Success = true
	resToken.Output = plainText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}
