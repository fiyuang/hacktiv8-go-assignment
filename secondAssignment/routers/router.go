package routers

import (
	"secondAssignment/controllers"
	"secondAssignment/db"

	"github.com/gin-gonic/gin"
)

func init() {
	db.InitializeDB()
}

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)

	return router
}
