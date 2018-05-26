package main

import(
	"github.com/kataras/iris"

	"gowork/muxibook/api"
)

func main(){
	muxibook:=iris.New()
	muxibook.Post("/signup",api.Signup)
	muxibook.Post("/login",api.Signin)

	muxibook.Post("/book",api.Book)
	muxibook.Post("/lendbook",api.BookLend)

	muxibook.Run(iris.Addr(":8080"))
}