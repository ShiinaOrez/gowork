package main

import (
	"github.com/ShiinaOrez/gowork/to-gather/api"
	"github.com/ShiinaOrez/gowork/to-gather/config"

	"github.com/kataras/iris"
	"github.com/ogier/pflag"
)

var (
	// like: go run main.go -c confg/config.yaml
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	toGather := iris.New()
	toGather.Post("/auth/login/", api.Login)
	//    toGather.Post("/activity/post/", api.LoginRequired(api.ActivityPost))
	toGather.Post("/activity/post/", api.LoginRequiredMiddleWare, api.ActivityPost)
	toGather.Get("/activity/{aid:int}/", api.LoginRequiredMiddleWare, api.ActivityGet)
	toGather.Post("/activity/{aid:int}/", api.LoginRequiredMiddleWare, api.ActivityPick)
	toGather.Put("/activity/{aid:int}/", api.LoginRequiredMiddleWare, api.ActivityPut)
	//    toGather.Get("/activity/pickable/list/{page:path}", api.ActivityPickableList)

	toGather.Run(iris.Addr(":8080"))
}
