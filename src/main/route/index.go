package route

import (
	"github.com/gin-gonic/gin"
)

// CreateRouter : create a router instance
func CreateRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	todoRouter := api.Group("/todo")

	setupTodoRoute(todoRouter)

	return r
}
