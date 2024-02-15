package model

type DecryptRequestToken struct {
	CipherText 	string	`json:"CipherText"`
	Key       	string	`json:"Key"`
}

type DecryptResponseToken struct {
	PlainText	string	`json:"PlainText"`
}
