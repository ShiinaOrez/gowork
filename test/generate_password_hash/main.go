package main

import (
	"fmt"
	"github.com/ShiinaOrez/GoSecurity/security"
)

func main() {
	password := "Muxi4ever"
	passwordHash := security.GeneratePasswordHash(password)
	fmt.Println(passwordHash)
	fmt.Println(security.CheckPasswordHash("mashiro", "pbkdf2:sha1:1000$wCMGTCep$38132208a6e553386702f084713a9eb8e497466a"))
}
