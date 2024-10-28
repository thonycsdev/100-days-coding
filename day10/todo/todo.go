package todo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Todo struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	Completed   bool      `json:"completed"`
	Due_date    time.Time `json:"due_date"`
}

func (t Todo) Print() {
	fmt.Printf("Title: %v\n", t.Title)
	fmt.Printf("Description: %v\n", t.Description)
	fmt.Printf("Completed: %v\n", t.Completed)
	fmt.Printf("Due Date: %v\n", t.Due_date.Format(time.DateTime))
	fmt.Printf("Create At: %v\n", t.Created_at.Format(time.DateTime))
	fmt.Println("=====================================================================")
}
func (t Todo) SaveTodo() {
	file, err := os.OpenFile("todos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicln(err)
	}
	jsonTodo := ParseToJson(&t)
	bytesWritten, err := fmt.Fprintln(file, jsonTodo)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Todo Saved!. Size: %v Bytes \n", bytesWritten)

}

func GetAllTodos() *[]Todo {
	var todos []Todo

	file, err := os.Open("todos.txt")
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var todo Todo
		json.Unmarshal(scanner.Bytes(), &todo)
		todos = append(todos, todo)
	}

	return &todos
}

func ParseToJson(todo *Todo) string {
	result, err := json.Marshal(*todo)
	if err != nil {
		log.Panicln(err)
	}
	return string(result)

}

func New(title, description string, due_date time.Time) *Todo {

	return &Todo{
		Title:       title,
		Due_date:    due_date,
		Description: description,
		Completed:   false,
		Created_at:  time.Now(),
	}

}
