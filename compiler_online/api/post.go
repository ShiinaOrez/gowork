package api

import (
	"github.com/ShiinaOrez/gowork/compiler_online/code"
	"github.com/ShiinaOrez/gowork/compiler_online/compiler"
	"github.com/ShiinaOrez/gowork/compiler_online/constant"
	"github.com/gin-gonic/gin"
)

type RunCodeRequest struct {
	TempID    int64 `json:"temp_id"`
	CodeBlock string `json:"code_block"`
}

func RunCode(c *gin.Context) {
	request := RunCodeRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"msg": "Bind post data failed. ",
			"code": constant.InvalidRequest,
		})
		return
	}
	tempStr := GetTempByID(request.TempID)
	codeObj := code.BuildCode(tempStr)
	codeObj.SetArgs(map[string]interface{}{
		"Code": request.CodeBlock,
	})
	compilerObj, err := compiler.NewCompilerWithCode(codeObj)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "Build compiler failed. ",
			"code": constant.ErrorInside,
		})
		return
	}
	output, err := compilerObj.Do()
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "Run code failed. ",
			"code": constant.CodeRunningFailed,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "ok",
		"code": constant.SuccessCode,
		"output": string(output),
	})
	return
}

func GetTempByID(tempID int64) string {
	return `
package main

{{ .Code }}

func main() {
    Do()
    return
}
`
}