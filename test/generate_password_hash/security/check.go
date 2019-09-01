package security

import (
	"strings"
)

func CheckPasswordHash(password string, hash string) bool {
	if strings.Count(hash, "$") < 2 {
		return false
	}
	pwdHashList := strings.Split(hash, "$")
	return pwdHashList[2] == hashInternal(pwdHashList[1], password)
}