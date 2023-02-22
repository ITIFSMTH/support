package adminController

import (
	"support-back/shared"

	"golang.org/x/crypto/bcrypt"
)

type AdminController struct{}

func getPasswordHash(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), shared.PW_COST)
	return string(bytes), err
}
