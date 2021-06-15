package app

import (
	"github.com/jgmc3012/bookstore_users-api/controllers/ping"
	"github.com/jgmc3012/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping/", ping.Ping)

	router.GET("/users/:user_id/", users.GetUser)
	router.GET("/users/search/", users.SearchUser)
	router.POST("/users/", users.CreateUser)
	router.PUT("/users/:user_id/", users.UpdateUser)

}
