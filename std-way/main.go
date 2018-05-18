package main

import(
	"fmt"
	"net/http"
	"github.com/kataras/iris"
)

func main(){
	app:=iris.New()

	app.Get("/",func (ctx iris.Context){
		ctx.Writef("hello from the server")
	})

	app.Get("/mypath",func (ctx iris.Context){
		ctx.HTML("<h1>hello!</h1>")
	})

	app.Build()

	srv:=&http.Server{Handler:app,Addr:":8080"}
	fmt.Printf("\n\n**welcome to shiina's test**\n")
	println("Start a server listening on http://localhost:8080")
	srv.ListenAndServe()
}