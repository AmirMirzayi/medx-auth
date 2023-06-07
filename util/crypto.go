package util

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwd), nil
}

func Verify(password, hash string) error {
return bcrypt.CompareHashAndPassword([]byte(password),[]byte(hash))
}
