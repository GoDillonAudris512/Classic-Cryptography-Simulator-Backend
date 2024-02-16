package middleware

import (
	"encoding/json"
	"net/http"

	"classic-crypt/model"
)

func BuildAutoKeyText(text string, key string) string {
	textLength := len(text)
	keyLength := len(key)

	if keyLength >= textLength {
		return key[:textLength]
	} else {
		keyText := key
		keyText += text[:(textLength - keyLength)]

		return keyText
	}
}

func HandleAutoVigenere(response http.ResponseWriter, request *http.Request) {
	var reqToken model.VigenereRequestToken
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqToken)

	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if reqToken.Encrypt {
		keyText := BuildAutoKeyText(reqToken.Input, reqToken.Key)
		EncryptVigenere(reqToken.Input, keyText, response)
	} else {
		DecryptAutoVigenere(reqToken.Input, reqToken.Key, response)
	}
}

func DecryptAutoVigenere(input string, key string, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	keyText := key
	plainText := ""
	for i := 0; i < len(input); i++ {
		token1 := alphabetToNumber[input[i]]
		token2 := alphabetToNumber[keyText[i]]

		var cipherToken uint8
		if token1-token2 < 0 {
			cipherToken = numberToAlphabet[(token1-token2+26)%26]
		} else {
			cipherToken = numberToAlphabet[(token1-token2)%26]
		}
		plainText += string(cipherToken)
		keyText += string(cipherToken)
	}

	var resToken model.VigenereResponseToken
	resToken.Success = true
	resToken.Output = plainText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}
