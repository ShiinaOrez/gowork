package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"time"
)

func main() {
	db := redis.New(service.Config{
		Network:     service.DefaultRedisNetwork,
		Addr:        service.DefaultRedisAddr,
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: service.DefaultRedisIdleTimeout,
		Prefix:      "",
	})

	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	defer db.Close()

	sess := sessions.New(sessions.Config{
		Cookie:  "sessionscookieid",
		Expires: 45 * time.Minute,
	})

	sess.UseDatabase(db)

	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("You should navigate to the /set, /get, /delete, /clear,/destroy instead")
	})

	app.Get("/set", func(ctx iris.Context) {
		s := sess.Start(ctx)
		s.Set("name", "iris")
		ctx.Writef("All ok session value of the 'name' is: %s", s.GetString("name"))
	})

	app.Get("/set/{key}/{value}", func(ctx iris.Context) {
		key, value := ctx.Params().Get("key"), ctx.Params().Get("value")
		s := sess.Start(ctx)
		s.Set(key, value)

		ctx.Writef("%s is %s", key, value)
	})

	app.Get("/get", func(ctx iris.Context) {
		name := sess.Start(ctx).GetString("name")

		ctx.Writef("name : %s", name)
	})

	app.Get("/get/{key}", func(ctx iris.Context) {
		name := sess.Start(ctx).GetString(ctx.Params().Get("key"))

		ctx.Writef("name : %s", name)
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
