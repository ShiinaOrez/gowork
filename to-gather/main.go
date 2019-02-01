package main

import (
    "github.com/ShiinaOrez/gowork/to-gather/api"
    "github.com/kataras/iris"
)

func main(){
    toGather := iris.New()
    toGather.Post("/auth/login/", api.Login)
//    toGather.Post("/activity/post/", api.LoginRequired(api.ActivityPost))
    toGather.Post("/activity/post/", api.LoginRequiredMiddleWare, api.ActivityPost)
    toGather.Get("/activity/{aid:int}/",api.LoginRequiredMiddleWare, api.ActivityGet)
    toGather.Post("/activity/{aid:int}/",api.LoginRequiredMiddleWare, api.ActivityPick)
    toGather.Put("/activity/{aid:int}/",api.LoginRequiredMiddleWare, api.ActivityPut)
//    toGather.Get("/activity/pickable/list/{page:path}", api.ActivityPickableList)

    toGather.Run(iris.Addr(":8080"))
}