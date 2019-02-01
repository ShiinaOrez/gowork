package api

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	_ "strconv"
	_ "strings"
)

func LoginRequired(function context.Handler) context.Handler {
	return func(ctx iris.Context) {
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
		function(ctx)
		return
	}
}