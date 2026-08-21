package pwd

import "golang.org/x/crypto/bcrypt"

func RawToHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareWithHash(raw, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
