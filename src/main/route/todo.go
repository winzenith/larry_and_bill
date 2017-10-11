package route

import (
	"github.com/gin-gonic/gin"

	"main/controller"
)

func setupTodoRoute(router *gin.RouterGroup) {
	todoController := &controller.Todo{}

	router.POST("/", todoController.Create)
	router.GET("/", todoController.List)
	router.DELETE("/", todoController.DeleteAll)
}
