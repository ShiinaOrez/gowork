package security

import (
	"fmt"
)

func GeneratePasswordHash(password string) string {
	salt := genSalt()
	hash := hashInternal(salt, password)
	return fmt.Sprintf("pbkdf2:sha1:%v$%s$%s", iterations, salt, hash)
}