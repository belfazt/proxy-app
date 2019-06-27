package main

import (
	handlers "github.com/belfazt/proxy-app/api/handlers"
	middleware "github.com/belfazt/proxy-app/api/middleware"
	server "github.com/belfazt/proxy-app/api/server"
	utils "github.com/belfazt/proxy-app/api/utils"
)

func main() {
	/*
		Router Iris
		ENV Vars
	*/

	utils.LoadEnv()
	middleware.Init()
	var app = server.SetUp()
	handlers.HandleRedirection(app)
	server.RunServer(app)
}
