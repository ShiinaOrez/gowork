package compiler

import (
	"errors"
	"fmt"
	"github.com/ShiinaOrez/gowork/compiler_online/code"
	"os"
	"os/exec"
)

type Compiler struct {
	hash       string
	content    string
}

func NewCompilerWithCode(c code.Code) (newCompiler Compiler, err error) {
	var (
		content string
	)
	content, err = c.Get()
	if err != nil {
		return
	}
	newCompiler.hash = c.GetHash()
	newCompiler.content = content
	return
}

func (c *Compiler) Do() (string, error) {
	var newFileName string = "gen-"+c.hash+".go"

	if err := func() error {
		file, err := os.Create(newFileName)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := file.WriteString(c.content); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return "", err
	}

	cmd := fmt.Sprintf("mkdir %s && mv %s %s/ && cd %s/ && go build -o main && ./main", c.hash, newFileName, c.hash, c.hash)
    out, err := exec.Command("bash", "-c", cmd).Output()
    defer exec.Command("rm", "-rf", c.hash)

    if err != nil {
    	return "", errors.New("exec cmd `"+cmd+"`"+" failed, reason: "+err.Error()+". ")
	}
    return string(out), nil
}
