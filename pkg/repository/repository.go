package repository

import "github.com/ProxS/todo-app"

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(login, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository interface {
	Autorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
