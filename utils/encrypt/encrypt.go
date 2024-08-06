package encrypt

import "golang.org/x/crypto/bcrypt"

func EncryptString(password string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(enc), nil
}
