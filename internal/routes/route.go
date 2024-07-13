package routes

import (
	"github.com/halowahyudi/rest-api-structure-golang/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", handlers.PingHandler)
    r.GET("/dbtest", handlers.DBTestHandler)
}
