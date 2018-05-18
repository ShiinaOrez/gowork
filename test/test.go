package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main(){
	shiina:=iris.New()
	shiina.Logger().SetLevel("debug")
	shiina.Use(recover.New())
	shiina.Use(logger.New())

	shiina.Handle("GET","/",func(ctx iris.Context){
		ctx.JSON(iris.Map{"message": "hello,my friend!"})
	})

	shiina.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed))
}
