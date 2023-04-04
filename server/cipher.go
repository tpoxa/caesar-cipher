package server

type Cipher interface {
	Encrypt(text string, shifts int) string
	Decrypt(text string, shifts int) string
}
