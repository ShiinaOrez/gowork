package main

import (
	"github.com/kataras/iris"
	"net/http"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("hello from the server")
	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("here is %s", ctx.Path())
	})

	srv := &http.Server{Addr: ":8080"}

	app.Run(iris.Server(srv))

}
