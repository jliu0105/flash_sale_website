package middleware

import "github.com/kataras/iris"

func AuthConProduct(ctx iris.Context) {

	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("must log in first")
		ctx.Redirect("/user/login")
		return
	}
	ctx.Application().Logger().Debug("already logged in")
	ctx.Next()
}
