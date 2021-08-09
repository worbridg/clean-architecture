package main

import (
	"github.com/worbridg/clean-architecture/di"
)

func main() {
	userHandler := di.UserHandler()
	router := di.Router()
	router.GET("/user", userHandler.Get)
	router.GET("/user/new", userHandler.Create)
	if err := router.Start(); err != nil {
		panic(err)
	}
}
