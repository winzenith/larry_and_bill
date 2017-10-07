package controller

import (
	"github.com/gin-gonic/gin"

	"main/database"
	"main/model"
)

var db = database.Get()

// Todo : todo controller
type Todo struct{}

// Create : create new todo
func (ctrl *Todo) Create(c *gin.Context) {
	content := c.PostForm("content")

	newTodo := &model.Todo{
		Content: content,
	}

	err := db.Insert(newTodo)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"ok": true,
	})
}

// List : list all todos
func (ctrl *Todo) List(c *gin.Context) {
	var todos []model.Todo
	err := db.Model(&todos).Select()

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": &todos,
	})
}

// DeleteAll : delete all items in table
func (ctrl *Todo) DeleteAll(c *gin.Context) {
	_, err := db.Model(&model.Todo{}).Where("TRUE").Delete()

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"ok": true,
	})
}
