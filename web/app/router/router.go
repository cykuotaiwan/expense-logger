package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"
var Router *gin.Engine

func Run() {
	fmt.Println("Starting at http://localhost" + PORT)
	Router.Run(PORT)
}
