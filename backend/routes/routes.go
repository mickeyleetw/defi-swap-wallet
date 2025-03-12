package routes

import (
	"defi-swap-backend/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/send", handlers.SendTransaction)

	return r
}
