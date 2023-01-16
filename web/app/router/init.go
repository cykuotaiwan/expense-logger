package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	mode := os.Getenv("OPMODE")
	if mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.Default()
	Router.GET("/", GetAlbums)
}
