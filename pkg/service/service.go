package service

import "github.com/ProxS/todo-app/pkg/repository"

type Autorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service interface {
	Autorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
