package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jmillandev/bookstore_utils-go/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("About to start the application")
	router.Run(":8080")
}
