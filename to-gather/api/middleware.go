package api

import "github.com/kataras/iris"

func LoginRequiredMiddleWare(ctx iris.Context) {
	token := Token(ctx.GetHeader("token"))
	statu, code := token.LoginRequired()
	if !statu {
		ctx.StatusCode(code)
		ctx.JSON(map[string]string{
			"msg": "authenticate fail",
		})
		return
	}
	ctx.Values().Set("uid", code)
	ctx.Next()
}
