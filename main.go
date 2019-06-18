package main

import (
	handlers "github.com/belfazt/proxy-app/api/handlers"
	server "github.com/belfazt/proxy-app/api/server"
	utils "github.com/belfazt/proxy-app/api/utils"
)

func main() {
	/*
		Router Iris
		ENV Vars
	*/

	utils.LoadEnv()

	handlers.HandleRedirection()

	server.RunServer(server.SetUp())
}
