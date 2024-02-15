package model

type EncryptRequestToken struct {
	PlainText string	`json:"PlainText"`
	Key       string	`json:"Key"`
}

type EncryptResponseToken struct {
	CipherText	string	`json:"CipherText"`
}
