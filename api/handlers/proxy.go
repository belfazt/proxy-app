package handlers

import (
	"github.com/kataras/iris"
)

// Should redirect traffic
func HandleRedirection(app *iris.Application) {
	app.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{"result": "ok"})
	})
}
