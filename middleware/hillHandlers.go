package middleware

import (
	"encoding/json"
	"net/http"

	"classic-crypt/lib"
	"classic-crypt/model"
)

func NormalizeMatrix(matrix *[][]int) {
	for i := 0; i < len(*matrix); i++ {
		for j := 0; j < len(*matrix); j++ {
			for (*matrix)[i][j] > 0 {
				(*matrix)[i][j] -= 26
			}

			for (*matrix)[i][j] < 0 {
				(*matrix)[i][j] += 26
			}
		}
	}
}

func ProcessHillInput(input string, m int) []string {
	processed := input
	for len(processed)%m != 0 {
		processed += "X"
	}

	var substrings []string
	for i := 0; i < len(processed); i += m {
		end := i + m
		substrings = append(substrings, processed[i:end])
	}

	return substrings
}

func MakeVector(substring string) [][]int {
	vector := make([][]int, len(substring))

	for i, char := range substring {
		vector[i] = []int{alphabetToNumber[uint8(char)]}
	}

	return vector
}

func HandleHill(response http.ResponseWriter, request *http.Request) {
	var reqToken model.HillRequestToken
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqToken)

	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	matrix := lib.Transpose(reqToken.Matrix)
	NormalizeMatrix(&matrix)

	det := lib.Determinant(matrix)
	for det < 0 {
		det += 26
	}

	gcd, x, _ := ExtendedEuclidean(det, 26)

	if gcd != 1 {
		response.Header().Set("Content-Type", "application/json")

		var resToken model.HillResponseToken
		resToken.Success = false
		resToken.Output = "Unable to process because the matrix does not have inverse modulo"

		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(resToken)
		return
	}

	substrings := ProcessHillInput(reqToken.Input, len(reqToken.Matrix))

	if reqToken.Encrypt {
		EncryptHill(substrings, &matrix, response)
	} else {
		detInv := (x + 26) % 26
		DecryptHill(substrings, &matrix, detInv, response)
	}
}

func EncryptHill(substrings []string, matrix *[][]int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	cipherText := ""
	for _, substring := range substrings {
		vector := MakeVector(substring)

		result := lib.Multiply(*matrix, vector)

		for i := 0; i < len(substring); i++ {
			token := result[i][0]
			cipherText += string(numberToAlphabet[token%26])
		}
	}

	var resToken model.HillResponseToken
	resToken.Success = true
	resToken.Output = cipherText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func DecryptHill(substrings []string, matrix *[][]int, detInv int, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	adjoint := lib.Adjoint(*matrix)
	inverse := lib.MultiplyWithConstant(adjoint, detInv)
	NormalizeMatrix(&inverse)

	plainText := ""
	for _, substring := range substrings {
		vector := MakeVector(substring)

		result := lib.Multiply(inverse, vector)

		for i := 0; i < len(substring); i++ {
			token := result[i][0]
			plainText += string(numberToAlphabet[token%26])
		}
	}

	var resToken model.HillResponseToken
	resToken.Success = true
	resToken.Output = plainText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}
