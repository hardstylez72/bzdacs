package user

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func hashAndSalt(pwd []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func Encrypt(stringToEncrypt string) (encryptedString string, err error) {
	return hashAndSalt([]byte(stringToEncrypt))
}
