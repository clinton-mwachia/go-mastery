package utils

import (
	"errors"
	"sync"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var (
	todos  = []Todo{}
	nextID = 1
	mu     sync.Mutex
)

func GetAll() []Todo {
	mu.Lock()
	defer mu.Unlock()

	todoCopy := make([]Todo, len(todos))
	copy(todoCopy, todos)
	return todoCopy
}

func Add(title string) {
	mu.Lock()
	defer mu.Unlock()

	todos = append(todos, Todo{
		ID:    nextID,
		Title: title,
		Done:  false,
	})
	nextID++
}

func Toggle(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = !todos[i].Done
			return nil
		}
	}
	return errors.New("todo item not found")
}

func Delete(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo item not found")
}
