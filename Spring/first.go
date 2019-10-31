package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type Data struct {
	Msg string `json:"msg"`
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("Hello, MUXI.")
}

func (this *MainController) Post() {
	var requestData Data
	var testData Data
	testData.Msg = "hello"
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, requestData)
	if err != nil {
		this.Ctx.Output.SetStatus(403)
		this.Ctx.WriteString(err.Error())
	}
	this.Data["json"] = &testData
	this.ServeJSON()
}

func main() {
	beego.Config{""}
	beego.Router("/", &MainController{})
	beego.Router("/post/", &MainController{})
	beego.Run()
}
