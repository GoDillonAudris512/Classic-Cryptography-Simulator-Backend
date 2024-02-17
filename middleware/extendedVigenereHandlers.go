package middleware

import (
	"encoding/json"
	"net/http"

	"classic-crypt/model"
)

func BuildExtendedKeyText(text []uint8, key []uint8) []uint8 {
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
		EncryptExtendedVigenere(reqToken.Input, keyText, response)
	} else {
		DecryptExtendedVigenere(reqToken.Input, keyText, response)
	}
}

func EncryptExtendedVigenere(input []uint8, keyText []uint8, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	cipherText := []uint8{}
	for i := 0; i < len(input); i++ {
		token1 := alphabet256ToNumber[input[i]]
		token2 := alphabet256ToNumber[keyText[i]]

		cipherToken := numberToAlphabet256[(token1+token2)%256]
		cipherText = append(cipherText, cipherToken)
	}

	var resToken model.ExtendedVigenereResponseToken
	resToken.Success = true
	resToken.Output = cipherText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func DecryptExtendedVigenere(input []uint8, keyText []uint8, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	plainText := []uint8{}
	for i := 0; i < len(input); i++ {
		token1 := alphabet256ToNumber[input[i]]
		token2 := alphabet256ToNumber[keyText[i]]

		plainToken := numberToAlphabet256[(token1-token2+256)%256]
		plainText = append(plainText, plainToken)
	}

	var resToken model.ExtendedVigenereResponseToken
	resToken.Success = true
	resToken.Output = plainText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}
