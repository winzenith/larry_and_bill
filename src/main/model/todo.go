package model

import "fmt"

// Todo : todo model
type Todo struct {
	ID      int64
	Content string `sql:",unique"`
}

func (t Todo) String() string {
	return fmt.Sprintf("Todo<%s>", t.Content)
}
