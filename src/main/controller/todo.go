package controller

import (
	"github.com/gin-gonic/gin"

	"main/db"
	"main/model"
)

var conn = db.GetConn()

// Todo : todo controller
type Todo struct{}

// Create : create new todo
func (ctrl *Todo) Create(c *gin.Context) {
	content := c.PostForm("content")

	_, err := conn.Exec("INSERT INTO todos(content) VALUES($1)", content)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"ok": true,
	})
}

// List : list all todos
func (ctrl *Todo) List(c *gin.Context) {
	rows, err := conn.Query("SELECT id, content FROM todos")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var id int64
		var content string

		err := rows.Scan(&id, &content)

		if err != nil {
			panic(err)
		}

		todo := model.Todo{Content: content, ID: id}
		todos = append(todos, todo)
	}

	c.JSON(200, gin.H{
		"data": &todos,
	})
}

// DeleteAll : delete all items in table
func (ctrl *Todo) DeleteAll(c *gin.Context) {
	_, err := conn.Exec("DELETE FROM todos")

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"ok": true,
	})
}
