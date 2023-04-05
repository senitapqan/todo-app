package service

import (
	"webapp/model"
	"webapp/pkg/repository"
)

type todoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *todoListService {
	return &todoListService{repo: repo}
}

// in process...
func (r *todoListService) Create(userId int, list model.TodoList) (int, error) {
	return r.repo.Create(userId, list)
}

// in process...
func (r *todoListService) GetAll(userId int) ([]model.TodoList, error) {
	return r.repo.GetAll(userId)
}

// in process...
func (r *todoListService) GetById(userId, listId int) (model.TodoList, error) {
	return r.repo.GetById(userId, listId)
}

// in process...
func (r *todoListService) Delete(userId, listId int) error{
	return r.repo.Delete(userId, listId)
}
