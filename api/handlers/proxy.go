package handlers

import (
	"github.com/belfazt/proxy-app/api/middleware"
	"github.com/kataras/iris"
)

// Should redirect traffic
func HandleRedirection(app *iris.Application) {
	app.Get("/", middleware.Handler, pingHandler)
}

func pingHandler(c iris.Context) {
	c.JSON(iris.Map{"result": "ok"})
}
