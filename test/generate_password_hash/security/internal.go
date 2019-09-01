package security

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

func hashInternal(salt, password string) string {
	hash := pbkdf2.Key([]byte(password), []byte(salt), iterations, 20, sha1.New)
	return hex.EncodeToString(hash)
}

func genSalt() string {
	var bytes = make([]byte, salt_length)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = SALT_CHARS[v%byte(len(SALT_CHARS))]
	}
	return string(bytes)
}