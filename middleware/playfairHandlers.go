package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"classic-crypt/model"
)

func ProcessKeyword(keyword string) string {
	processed := ""

	for _, char := range keyword {
		if char != 'J' && !strings.ContainsRune(processed, char) {
			processed += string(char)
		}
	}

	for i := 'A'; i <= 'Z'; i++ {
		if i != 'J' && !strings.ContainsRune(processed, i) {
			processed += string(i)
		}
	}

	return processed
}

func KeywordToMapPair(keyword string) model.MapPair {
	var mapPair model.MapPair
	mapPair.CharToPair = make(map[rune]model.Pair)
	mapPair.PairToChar = make(map[model.Pair]rune)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			index := i*5 + j
			char := rune(keyword[index])
			mapPair.CharToPair[char] = model.Pair{Row: i, Col: j}
			mapPair.PairToChar[model.Pair{Row: i, Col: j}] = char
		}
	}

	return mapPair
}

func ProcessPlayfairInput(input string) []model.Bigram {
	processed := strings.Map(func(r rune) rune {
		if r == 'J' {
			return 'I'
		}
		return r
	}, input)

	var bigrams []model.Bigram
	var bigram model.Bigram

	for _, char := range processed {
		if bigram.First == 0 {
			bigram.First = char
		} else {
			if char != bigram.First || char == 'X' {
				bigram.Second = char
				bigrams = append(bigrams, bigram)
				bigram = model.Bigram{}
			} else {
				bigram.Second = 'X'
				bigrams = append(bigrams, bigram)
				bigram = model.Bigram{}
				bigram.First = char
			}
		}
	}

	if bigram.First != 0 && bigram.Second == 0 {
		bigram.Second = 'X'
		bigrams = append(bigrams, bigram)
	}

	return bigrams
}

func HandlePlayfair(response http.ResponseWriter, request *http.Request) {
	var reqToken model.PlayfairRequestToken
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&reqToken)

	if err != nil {
		http.Error(response, "Invalid request payload", http.StatusBadRequest)
		return
	}

	key := ProcessKeyword(reqToken.Keyword)

	if reqToken.Encrypt {
		EncryptPlayfair(reqToken.Input, key, response)
	} else {
		DecryptPlayfair(reqToken.Input, key, response)
	}
}

func EncryptPlayfair(input string, key string, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	mapPair := KeywordToMapPair(key)
	bigrams := ProcessPlayfairInput(input)

	cipherText := ""
	for _, bigram := range bigrams {
		firstLoc := mapPair.CharToPair[bigram.First]
		secondLoc := mapPair.CharToPair[bigram.Second]

		if firstLoc.Row == secondLoc.Row {
			newFirstCol := (firstLoc.Col + 1) % 5
			newSecondCol := (secondLoc.Col + 1) % 5

			cipherText += string(mapPair.PairToChar[model.Pair{Row: firstLoc.Row, Col: newFirstCol}])
			cipherText += string(mapPair.PairToChar[model.Pair{Row: secondLoc.Row, Col: newSecondCol}])
		} else if firstLoc.Col == secondLoc.Col {
			newFirstRow := (firstLoc.Row + 1) % 5
			newSecondRow := (secondLoc.Row + 1) % 5

			cipherText += string(mapPair.PairToChar[model.Pair{Row: newFirstRow, Col: firstLoc.Col}])
			cipherText += string(mapPair.PairToChar[model.Pair{Row: newSecondRow, Col: secondLoc.Col}])
		} else {
			cipherText += string(mapPair.PairToChar[model.Pair{Row: firstLoc.Row, Col: secondLoc.Col}])
			cipherText += string(mapPair.PairToChar[model.Pair{Row: secondLoc.Row, Col: firstLoc.Col}])
		}
	}

	var resToken model.PlayfairResponseToken
	resToken.Success = true
	resToken.Output = cipherText
	resToken.Key = key

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}

func DecryptPlayfair(input string, key string, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")

	mapPair := KeywordToMapPair(key)
	bigrams := ProcessPlayfairInput(input)

	plainText := ""
	for _, bigram := range bigrams {
		firstLoc := mapPair.CharToPair[bigram.First]
		secondLoc := mapPair.CharToPair[bigram.Second]

		if firstLoc.Row == secondLoc.Row {
			newFirstCol := (firstLoc.Col + 4) % 5
			newSecondCol := (secondLoc.Col + 4) % 5

			plainText += string(mapPair.PairToChar[model.Pair{Row: firstLoc.Row, Col: newFirstCol}])
			plainText += string(mapPair.PairToChar[model.Pair{Row: secondLoc.Row, Col: newSecondCol}])
		} else if firstLoc.Col == secondLoc.Col {
			newFirstRow := (firstLoc.Row + 4) % 5
			newSecondRow := (secondLoc.Row + 4) % 5

			plainText += string(mapPair.PairToChar[model.Pair{Row: newFirstRow, Col: firstLoc.Col}])
			plainText += string(mapPair.PairToChar[model.Pair{Row: newSecondRow, Col: secondLoc.Col}])
		} else {
			plainText += string(mapPair.PairToChar[model.Pair{Row: firstLoc.Row, Col: secondLoc.Col}])
			plainText += string(mapPair.PairToChar[model.Pair{Row: secondLoc.Row, Col: firstLoc.Col}])
		}
	}

	var resToken model.PlayfairResponseToken
	resToken.Success = true
	resToken.Output = plainText
	resToken.Key = key

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(resToken)
}