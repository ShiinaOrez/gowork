package code

import (
	"bytes"
	"github.com/ShiinaOrez/gowork/compiler_online/utils"
	"text/template"
)

type Code struct {
	hash  string
	temp  string
	args  map[string]interface{}
}

func BuildCode(temp string) Code {
	code := Code{}
	code.Init(temp)
	return code
}

func (c *Code) Init(temp string) {
	c.temp = temp
	c.hash = utils.NewHash(10)
}

func (c *Code) GetHash() string {
	return c.hash
}

func (c *Code) SetArgs(args map[string]interface{}) {
	c.args = args
}

func (c *Code) Get() (string, error) {
	var (
		result string
		output bytes.Buffer
		temp   *template.Template
		err    error
	)

	if temp, err = template.New(c.hash).Parse(c.temp); err != nil {
		return "", err
	}

	if err = temp.Execute(&output, c.args); err != nil {
		return "", err
	}

	result = output.String()

	return result, nil
}