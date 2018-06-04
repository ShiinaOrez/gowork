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
    muxibook.Post("/returnbook",api.BookReturn)
    muxibook.Post("/renewbook",api.BookRenew)
    muxibook.Post("/mybook",api.BookMy)
	muxibook.Post("/searchbook",api.Search)
	muxibook.Get("/getbook",api.BookGet)

	muxibook.Run(iris.Addr(":8080"))
}