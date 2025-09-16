package hashpassword

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

    if err != nil {
        return "", err
    }

    return string(hashPass), nil
}

func IsPasswordMatch(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

    return err == nil
}