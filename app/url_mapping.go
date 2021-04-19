package app

import "github.com/aushafy/go-bookstore-users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
