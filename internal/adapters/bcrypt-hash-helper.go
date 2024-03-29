package adapters

import "golang.org/x/crypto/bcrypt"

type BcryptHashHelper struct{}

func NewBcryptHashHelper() *BcryptHashHelper {
	return &BcryptHashHelper{}
}

func (a *BcryptHashHelper) GenerateHash(password string, salt int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err	
}

func (a *BcryptHashHelper) CompareHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}