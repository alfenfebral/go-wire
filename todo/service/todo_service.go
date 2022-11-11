package service

import (
	"go-clean-architecture/todo/models"
	"go-clean-architecture/todo/repository"
)

// TodoService represent the todo service
type TodoService interface {
	GetAll(keyword string, limit int, offset int) ([]*models.Todo, int, error)
	GetByID(id string) (*models.Todo, error)
	Create(value *models.Todo) (*models.Todo, error)
	Update(id string, value *models.Todo) (*models.Todo, error)
	Delete(id string) error
}

type TodoServiceImpl struct {
	todoRepo repository.MongoTodoRepository
}

// NewTodoService will create new an TodoService object representation of TodoService interface
func NewTodoService(a repository.MongoTodoRepository) TodoService {
	return &TodoServiceImpl{
		todoRepo: a,
	}
}

// GetAll - get all todo service
func (a *TodoServiceImpl) GetAll(keyword string, limit int, offset int) ([]*models.Todo, int, error) {
	res, err := a.todoRepo.FindAll(keyword, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Count total
	total, err := a.todoRepo.CountFindAll(keyword)
	if err != nil {
		return nil, 0, err
	}

	return res, total, nil
}

// GetByID - get todo by id service
func (a *TodoServiceImpl) GetByID(id string) (*models.Todo, error) {
	res, err := a.todoRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create - creating todo service
func (a *TodoServiceImpl) Create(value *models.Todo) (*models.Todo, error) {
	res, err := a.todoRepo.Store(&models.Todo{
		Title:       value.Title,
		Description: value.Description,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update - update todo service
func (a *TodoServiceImpl) Update(id string, value *models.Todo) (*models.Todo, error) {
	_, err := a.todoRepo.CountFindByID(id)
	if err != nil {
		return nil, err
	}

	_, err = a.todoRepo.Update(id, &models.Todo{
		Title:       value.Title,
		Description: value.Description,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete - delete todo service
func (a *TodoServiceImpl) Delete(id string) error {
	err := a.todoRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
