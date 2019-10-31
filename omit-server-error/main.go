package main

import (
	"github.com/kataras/iris"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
	return
}

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>hello,iris</h1>")
	})

	err := app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	check(err)
}
