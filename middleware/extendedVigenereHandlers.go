package middleware

import (
	"encoding/json"
	"net/http"

	"classic-crypt/model"
)

func BuildExtendedKeyText(text []int, key []int) []int {
	textLength := len(text)
	keyLength := len(key)

	if keyLength >= textLength {
		return key[:textLength]
	} else {
		keyText := key

		for len(keyText) < textLength {
			keyText = append(keyText, key...)
		}

		return keyText[:textLength]
	}
}

func HandleExtendedVigenere(response http.ResponseWriter, request *http.Request) {
	var reqToken model.ExtendedVigenereRequestToken
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqToken)

	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	keyText := BuildExtendedKeyText(reqToken.Input, reqToken.Key)

	if reqToken.Encrypt {
		HandleEncryptExtendedVigenere(reqToken.Input, keyText, response)
	} else {
		HandleDecryptExtendedVigenere(reqToken.Input, keyText, response)
	}
}

func HandleEncryptExtendedVigenere(input []int, keyText []int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	cipherText := EncryptExtendedVigenere(input, keyText)

	var resToken model.ExtendedVigenereResponseToken
	resToken.Success = true
	resToken.Output = cipherText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func EncryptExtendedVigenere(input []int, keyText []int) []int {
	cipherText := []int{}

	for i := 0; i < len(input); i++ {
		token1 := alphabet256ToNumber[uint8(input[i])]
		token2 := alphabet256ToNumber[uint8(keyText[i])]

		cipherToken := numberToAlphabet256[(token1+token2)%256]
		cipherText = append(cipherText, int(cipherToken))
	}

	return cipherText
}

func HandleDecryptExtendedVigenere(input []int, keyText []int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	plainText := DecryptExtendedVigenere(input, keyText)

	var resToken model.ExtendedVigenereResponseToken
	resToken.Success = true
	resToken.Output = plainText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func DecryptExtendedVigenere(input []int, keyText []int) []int {
	plainText := []int{}

	for i := 0; i < len(input); i++ {
		token1 := alphabet256ToNumber[uint8(input[i])]
		token2 := alphabet256ToNumber[uint8(keyText[i])]

		plainToken := numberToAlphabet256[(token1-token2+256)%256]
		plainText = append(plainText, int(plainToken))
	}

	return plainText
}