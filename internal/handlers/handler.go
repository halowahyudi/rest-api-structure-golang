package handlers

import (
	"net/http"

	"github.com/halowahyudi/rest-api-structure-golang/internal/db"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}

func DBTestHandler(c *gin.Context) {
    var version string
    err := db.DB.QueryRow("SELECT VERSION()").Scan(&version)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "version": version,
    })
}
