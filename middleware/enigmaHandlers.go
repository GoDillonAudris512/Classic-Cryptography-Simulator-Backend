package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"classic-crypt/model"
)

func IndexOfElement(array []uint8, target uint8) int {
	for i, num := range array {
		if num == target {
			return i
		}
	}
	return -1
}

func HandlePlugboardError(response http.ResponseWriter, error string) {
	var resToken model.EnigmaResponseToken
	resToken.Success = false
	resToken.Output = error

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func BuildPlugboard(plugboardText string, response http.ResponseWriter) {
	substrings := strings.Split(plugboardText, " ")
	if substrings[len(substrings)-1] == "" {
		substrings = substrings[:len(substrings)-1]
	}

	for index, substring := range substrings {
		if index == 10 {
			HandlePlugboardError(response, "Unable to process because Enigma Cipher only supports ten pairs of plug in plugboard")
		}
		if len(substring) != 2 {
			HandlePlugboardError(response, "Unable to process because a substring in plugboard has more than 2 alphabets")
		}

		first := substring[0]
		second := substring[1]

		if _, ok1 := plugboard[uint8(alphabetToNumber[first])]; ok1 {
			HandlePlugboardError(response, "Unable to process because an alphabet has 2 pairs in plugboard")
		}
		if _, ok2 := plugboard[uint8(alphabetToNumber[second])]; ok2 {
			HandlePlugboardError(response, "Unable to process because an alphabet has 2 pairs in plugboard")
		}

		plugboard[uint8(alphabetToNumber[first])] = uint8(alphabetToNumber[second])
		plugboard[uint8(alphabetToNumber[second])] = uint8(alphabetToNumber[first])
	}

	for i := uint8(0); i < uint8(26); i++ {
		if _, ok := plugboard[i]; !ok {
			plugboard[i] = i
		}
	}
}

func BuildEnigmaConfig(reqToken model.EnigmaRequestToken, response http.ResponseWriter) {
	rotorOrder[reqToken.Order1-1] = 0
	rotorOrder[reqToken.Order2-1] = 1
	rotorOrder[reqToken.Order3-1] = 2
	pos[0] = reqToken.Position1 - 1
	pos[1] = reqToken.Position2 - 1
	pos[2] = reqToken.Position3 - 1
	BuildPlugboard(reqToken.Plugboard, response)
}

func PlugboardSwap(token uint8) uint8 {
	return plugboard[token]
}

func RotateRotors() {
	if pos[rotorOrder[2]] == turnover[2] {
		pos[rotorOrder[1]] = (pos[rotorOrder[1]] + 1) % 26
	}

	if pos[rotorOrder[1]] == turnover[1] {
		pos[rotorOrder[0]] = (pos[rotorOrder[0]] + 1) % 26
	}

	pos[rotorOrder[2]] = (pos[rotorOrder[2]] + 1) % 26
}

func ForwardSubstitution(token uint8) uint8 {
	newToken := token

	for i := 2; i >= 0; i-- {
		newToken = (newToken + uint8(pos[rotorOrder[i]])) % 26
		newToken = rotors[rotorOrder[i]][newToken]
		newToken = (newToken - uint8(pos[rotorOrder[i]]) + 26) % 26
	}

	return newToken
}

func Reflect(token uint8) uint8 {
	return reflector[token]
}

func BackwardSubstitution(token uint8) uint8 {
	newToken := token

	for i := 0; i <= 2; i++ {
		newToken = (newToken + uint8(pos[rotorOrder[i]])) % 26
		newToken = (uint8(IndexOfElement(rotors[rotorOrder[i]], newToken)) - uint8(pos[rotorOrder[i]]) + 26) % 26
	}

	return newToken
}

func ResetEnigmaConfig() {
	rotorOrder = make([]int, 3)
	pos = make([]int, 3)
	plugboard = make(map[uint8]uint8)
}

func EnigmaCipher(token uint8) uint8 {
	processedToken := token

	processedToken = PlugboardSwap(processedToken)
	RotateRotors()
	processedToken = ForwardSubstitution(processedToken)
	processedToken = Reflect(processedToken)
	processedToken = BackwardSubstitution(processedToken)
	processedToken = PlugboardSwap(processedToken)

	return processedToken
}

func HandleEnigma(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var reqToken model.EnigmaRequestToken
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqToken)

	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	BuildEnigmaConfig(reqToken, response)

	processedText := ""
	for _, char := range reqToken.Input {
		processedText += string(numberToAlphabet[int(EnigmaCipher(uint8(alphabetToNumber[uint8(char)])))])
	}

	ResetEnigmaConfig()

	var resToken model.EnigmaResponseToken
	resToken.Success = true
	resToken.Output = processedText

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}
