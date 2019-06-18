package server

import (
	"fmt"
	"github.com/kataras/iris"
	"os"
)

func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel(os.Getenv("DEBUG_LEVEL"))
	return app
}

func RunServer(app *iris.Application) {
	app.Run(
		iris.Addr(fmt.Sprintf(":%s", os.Getenv("PORT"))),
	)
}
