package route

import (
	"github.com/gin-gonic/gin"

	"main/controller"
)

func setupTodoRoute(router *gin.RouterGroup) {
	todoController := &controller.Todo{}

	router.GET("/", todoController.List)
	router.POST("/", todoController.Create)
	router.DELETE("/", todoController.DeleteAll)
}
