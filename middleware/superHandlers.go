package middleware

import (
	"encoding/json"
	"net/http"

	"classic-crypt/model"
)

func HandleSuper(response http.ResponseWriter, request *http.Request) {
	var reqToken model.SuperRequestToken
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqToken)

	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	keyText := BuildExtendedKeyText(reqToken.Input, reqToken.Key1)

	if reqToken.Encrypt {
		EncryptSuper(reqToken.Input, keyText, reqToken.Key2, response)
	} else {
		DecryptSuper(reqToken.Input, keyText, reqToken.Key2, response)
	}
}

func EncryptSuper(input []int, keyText []int, order int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	cipherText := EncryptExtendedVigenere(input, keyText)

	processedCipherText := []int{}
	for i := 0; i < order; i++ {
		for j := 0; j < len(cipherText); j++ {
			if j % order == i {
				processedCipherText = append(processedCipherText, cipherText[j])
			}
		}
	}

	var resToken model.SuperResponseToken
	resToken.Success = true
	resToken.Output = processedCipherText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func DecryptSuper(input []int, keyText []int, order int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	cap := len(input) / order
	copy := input
	numbers := make([][]int, order)
	for i := 0; i < order; i++ {
		if i < len(input) % order {
			numbers[i] = copy[:(cap+1)]
			copy = copy[(cap+1):]
		} else {
			numbers[i] = copy[:(cap)]
			copy = copy[(cap):]
		}
	}

	processedPlainText := []int{}
	column := 0
	for i := 0; i < len(input); i++ {
		if (len(numbers[column]) > 0) {
			processedPlainText = append(processedPlainText, numbers[column][0])
			numbers[column] = numbers[column][1:]
			column = (column + 1) % order
		}
	}

	plainText := DecryptExtendedVigenere(processedPlainText, keyText)

	var resToken model.SuperResponseToken
	resToken.Success = true
	resToken.Output = plainText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}