package middleware

import (
	"encoding/json"
	"net/http"
	"sort"
	"math"

	"classic-crypt/model"
)

func BuildColumnarKeyOrder(key []uint8) []int {
	sorted := key
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	order := []int{}
	first := true
	prevChar := uint8(0)
	for _, char := range sorted {
		if first {
			for i, keyToken := range key {
				if keyToken == char {
					order = append(order, i)
				}
			}
			prevChar = char
			first = false
		} else {
			if char != prevChar {
				for i, keyToken := range key {
					if keyToken == char {
						order = append(order, i)
					}
				}
				prevChar = char
			}
		}
	}

	return order
}

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
	order := BuildColumnarKeyOrder(reqToken.Key2)

	if reqToken.Encrypt {
		EncryptSuper(reqToken.Input, keyText, order, response)
	} else {
		DecryptSuper(reqToken.Input, keyText, order, response)
	}
}

func EncryptSuper(input []uint8, keyText []uint8, order []int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	cipherText := EncryptExtendedVigenere(input, keyText)

	processedCipherText := []uint8{}
	for _, rank := range order {
		for i, char := range cipherText {
			if i % len(order) == rank {
				processedCipherText = append(processedCipherText, char)
			}
		}
	}

	var resToken model.SuperResponseToken
	resToken.Success = true
	resToken.Output = processedCipherText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func DecryptSuper(input []uint8, keyText []uint8, order []int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	row := int(math.Ceil(float64(len(input)) / float64(len(order))))
	matrix := make([][]uint8, row)
	for i := 0; i < row - 1; i++ {
		matrix[i] = make([]uint8, len(order))
	}
	matrix[row - 1] = make([]uint8, len(input) % len(order))
	
	index := 0
	for _, rank := range order {
		colSize := 0
		if rank < len(input) % len(order) {
			colSize = row
		} else {
			colSize = row - 1
		}

		for i := 0; i < colSize; i++ {
			matrix[i][rank] = input[index]
			index++
		}
	}

	processedPlainText := []uint8{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			processedPlainText = append(processedPlainText, matrix[i][j])
		}
	}

	plainText := DecryptExtendedVigenere(processedPlainText, keyText)

	var resToken model.SuperResponseToken
	resToken.Success = true
	resToken.Output = plainText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}