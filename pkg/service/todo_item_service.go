package service

import (
	"webapp/model"
	"webapp/pkg/repository"
)

type todoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem) *todoItemService {
	return &todoItemService{repo: repo}
}

//in process...
func (r *todoItemService) Create(userId, listId int, item model.TodoItem) (int, error) {
	_, err := r.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return r.repo.Create(userId, listId, item)
}

//in process...
func (r *todoItemService) GetAll(userId, listId int) ([]model.TodoItem, error) {
	_, err := r.listRepo.GetById(userId, listId)
	if err != nil {
		return nil, err
	}

	return r.repo.GetAll(userId, listId)	
}

//in process...
func (r *todoItemService) GetById(userId, itemId int) (model.TodoItem, error) {
	return r.repo.GetById(userId, itemId)
}

//in process...
func (r *todoItemService) Delete(userId, listId, itemId int) error {
	return nil
}