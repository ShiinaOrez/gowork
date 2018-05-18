package main

import (
	"net/url"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/host"
)

func main(){
	app:=iris.New()
	app.Get("/",func (ctx iris.Context){
		ctx.Writef("hello from shiina_orez")
	})
	app.Get("/mypath/",func (ctx iris.Context){
		ctx.Writef("hello from S?O on path --/mypath")
	})

	target,_:=url.Parse("https://127.0.0.1:443")
	go host.NewProxy("127.0.0.1:80",target).ListenAndServe()

	app.Run(iris.TLS("127.0.0.1:443","mycert.cert","mykey.key"))
}