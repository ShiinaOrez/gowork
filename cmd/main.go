package main

import (
    "fmt"
    "github.com/ShiinaOrez/gowork/compiler_online/utils"
    "os"
    "os/exec"
)

const content = `
package main

import "fmt"

func main() {
    fmt.Println("hello, world!")
}
`

func main() {
    hash := utils.NewHash(10)
    file, err := os.Create(hash+".go")
    if err != nil {
        panic(err)
    }
    file.WriteString(content)
    file.Close()

    cmd := fmt.Sprintf("mkdir %s && mv %s.go %s/ && cd %s/ && pwd && go build -o main && ./main && cd .. && rm -rf %s/", hash, hash, hash, hash, hash)
    out, err := exec.Command("bash", "-c", cmd).Output()
    if err != nil {
        panic(err)
    }
    fmt.Print(string(out))
}

