package app

import (
	"github.com/jgmc3012/bookstore_users-api/controllers/ping"
	"github.com/jgmc3012/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping/", ping.Ping)

	router.POST("/v1/users/", users.Create)
	router.GET("/v1/users/:user_id/", users.Get)
	router.PUT("/v1/users/:user_id/", users.Update)
	router.PATCH("/v1/users/:user_id/", users.Update)
	router.DELETE("/v1/users/:user_id/", users.Delete)
	router.POST("/v1/login/", users.Login)
	router.GET("/v1/internal/users/search/", users.Search)
}
