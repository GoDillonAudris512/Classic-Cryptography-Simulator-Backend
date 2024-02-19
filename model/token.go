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
	Input   []int `json:"input"`
	Key     []int `json:"key"`
	Encrypt bool   	`json:"encrypt"`
}

type ExtendedVigenereResponseToken struct {
	Success bool   	`json:"success"`
	Output  []int `json:"output"`
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

// For Hill Cipher
type HillRequestToken struct {
	Input   string 	`json:"input"`
	Matrix	[][]int	`json:"matrix"`
	Encrypt bool   	`json:"encrypt"`
}

type HillResponseToken struct {
	Success bool   	`json:"success"`
	Output  string 	`json:"output"`
}

// For Super Encryption
type SuperRequestToken struct {
	Input   []int `json:"input"`
	Key1    []int `json:"key1"`
	Key2    int		`json:"key2"`
	Encrypt bool    `json:"encrypt"`
}

type SuperResponseToken struct {
	Success bool    `json:"success"`
	Output  []int `json:"output"`
}

// For Enigma Cipher
type EnigmaRequestToken struct {
	Input   	string 	`json:"input"`
	Position1	int		`json:"pos1"`
	Position2	int		`json:"pos2"`
	Position3	int 	`json:"pos3"`
	Order1		int		`json:"order1"`
	Order2		int		`json:"order2"`
	Order3		int		`json:"order3"`
	Plugboard	string	`json:"plugboard"`
}

type EnigmaResponseToken struct {
	Success bool   	`json:"success"`
	Output  string 	`json:"output"`
}