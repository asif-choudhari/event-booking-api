package utils

import "golang.org/x/crypto/bcrypt"

func Encode(data string) (string, error) {
	encodedData, err := bcrypt.GenerateFromPassword([]byte(data), 14)
	if err != nil {
		return "", err
	}

	return string(encodedData), nil
}

func Compare(data string, encodedData string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedData), []byte(data))
	return err == nil
}
