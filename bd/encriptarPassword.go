package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(passw string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(passw), costo)

	return string(bytes), err
}
