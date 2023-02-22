package authController

import (
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func comparePasswordHash(p string, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
