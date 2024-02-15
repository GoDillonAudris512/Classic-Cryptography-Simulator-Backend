package middleware

import (
	"encoding/base64"
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

func EncryptVigenere(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var encryptReqToken model.EncryptRequestToken
	err := json.NewDecoder(request.Body).Decode(&encryptReqToken)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	keyText := BuildKeyText(encryptReqToken.PlainText, encryptReqToken.Key)

	cipherText := ""
	for i := 0; i < len(encryptReqToken.PlainText); i++ {
		token1 := alphabetToNumber[encryptReqToken.PlainText[i]]
		token2 := alphabetToNumber[keyText[i]]
		
		cipherToken := numberToAlphabet[(token1+token2)%26]
		cipherText += string(cipherToken)
	}

	var encryptResToken model.EncryptResponseToken
	encryptResToken.CipherText = base64.StdEncoding.EncodeToString([]byte(cipherText))

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(encryptResToken)
}

func DecryptVigenere(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var decryptReqToken model.DecryptRequestToken
	err := json.NewDecoder(request.Body).Decode(&decryptReqToken)
	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	keyText := BuildKeyText(decryptReqToken.CipherText, decryptReqToken.Key)

	plainText := ""
	for i := 0; i < len(decryptReqToken.CipherText); i++ {
		token1 := alphabetToNumber[decryptReqToken.CipherText[i]]
		token2 := alphabetToNumber[keyText[i]]

		var cipherToken uint8
		if token1-token2 < 0 {
			cipherToken = numberToAlphabet[(token1-token2+26)%26]
		} else {
			cipherToken = numberToAlphabet[(token1-token2)%26]
		}
		plainText += string(cipherToken)
	}

	var decryptResToken model.DecryptResponseToken
	decryptResToken.PlainText = base64.StdEncoding.EncodeToString([]byte(plainText))

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(decryptResToken)
}
