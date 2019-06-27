package handlers

import (
	"encoding/json"
	"github.com/belfazt/proxy-app/api/middleware"
	"github.com/kataras/iris"
)

// Should redirect traffic
func HandleRedirection(app *iris.Application) {
	app.Get("/", middleware.Handler, pingHandler)
}

func pingHandler(c iris.Context) {
	response, err := json.Marshal(middleware.Queue)

	if err != nil {
		c.JSON(iris.Map{"status": 400, "result": "parse error"})
		return
	}

	c.JSON(iris.Map{"result": string(response)})
}
