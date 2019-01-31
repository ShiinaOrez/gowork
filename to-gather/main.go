package main

import(
    "github.com/kataras/iris"
    "github.com/ShiinaOrez/gowork/to-gather/api"
)

func main(){
    toGather := iris.New()
    toGather.Post("/auth/login/", api.Login)
    toGather.Post("/activity/post/", api.ActivityPost)
    toGather.Get("/activity/{aid:int}/", api.ActivityGet)
    toGather.Post("/activity/{aid:int}/", api.ActivityPick)

    toGather.Run(iris.Addr(":8080"))
}