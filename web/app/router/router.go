package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"
var Router *gin.Engine

func Run() {
	fmt.Println("Starting at http://localhost" + PORT)
	Router.Run(PORT)
}

func Init() {
	mode := os.Getenv("OPMODE")
	if mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.Default()
	Router.GET("/", GetAlbums)
}
