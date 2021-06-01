package app

import (
	"github.com/jgmc3012/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
