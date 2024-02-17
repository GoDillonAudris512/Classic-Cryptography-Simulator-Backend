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

// For Extended Vigenere Cipher
type ExtendedVigenereRequestToken struct {
	Input   []uint8 `json:"input"`
	Key     []uint8 `json:"key"`
	Encrypt bool   	`json:"encrypt"`
}

type ExtendedVigenereResponseToken struct {
	Success bool   	`json:"success"`
	Output  []uint8 `json:"output"`
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

// For Affine Cipher
type AffineRequestToken struct {
	Input 		string	`json:"input"`
	Slope		int 	`json:"slope"`
	Intercept	int 	`json:"intercept"`
	Encrypt		bool   	`json:"encrypt"`
}

type AffineResponseToken struct {
	Success 	bool   	`json:"success"`
	Output  	string 	`json:"output"`
}