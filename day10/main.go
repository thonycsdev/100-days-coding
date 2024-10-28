package main

import (
	"go-challange/todo_list/todo"
)

func main() {

	todos := todo.GetAllTodos()
	for _, val := range *todos {
		val.Print()
	}
}
