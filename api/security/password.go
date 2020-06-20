package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash make a PassWord hash
func Hash(PassWord string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(PassWord), bcrypt.DefaultCost)
}

// VerifyPassWord verify the hashed PassWord
func VerifyPassWord(hashedPassWord, PassWord string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassWord), []byte(PassWord))
}
