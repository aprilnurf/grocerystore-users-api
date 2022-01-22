package app

import(
	"github.com/aprilnurf/grocerystore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)
func StartApplication(){
	mapUrls()
	logger.Info("Starting application..")
	router.Run(":9999")
}