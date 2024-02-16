package model

// For Vigenere Cipher and Auto-Key Vigenere Cipher
type VigenereRequestToken struct {
	Input   string 	`json:"input"`
	Key     string 	`json:"key"`
	Encrypt bool   	`json:"encrypt"`
}

type VigenereResponseToken struct {
	Success bool   	`json:"success"`
	Output  string 	`json:"output"`
}

// For Playfair Cipher
type PlayfairRequestToken struct {
	Input 	string	`json:"input"`
	Keyword	string 	`json:"keyword"`
	Encrypt	bool   	`json:"encrypt"`
}

type PlayfairResponseToken struct {
	Success bool   	`json:"success"`
	Output  string 	`json:"output"`
	Key		string	`json:"key"`
}
