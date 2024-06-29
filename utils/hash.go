package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	cost := 15 //Cost can also come from valut or secrets
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashPass), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
