package security

import "golang.org/x/crypto/bcrypt"

func Hash(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

func VerityPassword(hp string, p string) error {
	return bcrypt.CompareHashAndPassword([]byte(hp), []byte(p))
}
