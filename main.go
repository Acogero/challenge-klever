package main

// import "fmt"
import (
	"challenge/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", controller.HealthCheck)
	r.GET("/details/:address", controller.GetDetails)

	r.Run()
}
