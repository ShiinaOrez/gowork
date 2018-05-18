package main

import (
	"encoding/xml"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"fmt"
)

type Msg struct{
	Inc string `json:"inc"`
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       string `json:"age"`
}

type ExampleXML struct {
	XMLName xml.Name `xml:"example"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

func main() {
	app:=iris.New()
	app.Post("/decode/{name:string}",func (ctx iris.Context){
		var user User
		ctx.ReadJSON(&user)
		msg:=Msg{
			Inc: "successful!",
		}
		ctx.StatusCode(200)
		ctx.JSON(msg)
		fmt.Println(ctx.Params().Get("name"))
//		ctx.Writef("%s %s is %s years old and comes from %s!",user.Firstname,user.Lastname,user.Age,user.City)

	})

	app.Get("/encode",func (ctx iris.Context){
		peter:=User{
			Firstname:"John",
			Lastname: "Doe",
			City: "Neither FBI knows!!",
			Age: "25",
		}
		ctx.JSON(peter)
	})

	app.Get("/binary",func (ctx iris.Context){
		ctx.Binary([]byte("some binary data here"))
	})

	app.Get("/text",func (ctx iris.Context){
		ctx.Text("plain text here")
	})

	app.Get("/json",func (ctx iris.Context){
		ctx.JSON(map[string]string {"hello": "json"})
	})

	app.Get("/jsonp",func (ctx iris.Context){
		ctx.JSONP(map[string]string {"hello": "jsonp"},context.JSONP{Callback: "callbackname"})
	})

	app.Get("/xml",func (ctx iris.Context){
		ctx.XML(ExampleXML{One:"Hello",Two:"xml"})
	})

	app.Get("/markdown",func (ctx iris.Context){
		ctx.Markdown([]byte("# hello dynamic markdown -- iris"))
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}
