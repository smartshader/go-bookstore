package app

import (
	"smartshader/go-bookstore/users/controllers/ping"
	"smartshader/go-bookstore/users/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", users.GetUser)
	router.POST("/user", users.CreateUser)
}
