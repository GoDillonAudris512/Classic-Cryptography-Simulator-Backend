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