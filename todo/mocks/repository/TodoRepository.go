// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	models "go-clean-architecture/todo/models"

	mock "github.com/stretchr/testify/mock"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// CountFindAll provides a mock function with given fields: keyword
func (_m *TodoRepository) CountFindAll(keyword string) (int, error) {
	ret := _m.Called(keyword)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(keyword)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountFindByID provides a mock function with given fields: id
func (_m *TodoRepository) CountFindByID(id string) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *TodoRepository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: keyword, limit, offset
func (_m *TodoRepository) FindAll(keyword string, limit int, offset int) ([]*models.Todo, error) {
	ret := _m.Called(keyword, limit, offset)

	var r0 []*models.Todo
	if rf, ok := ret.Get(0).(func(string, int, int) []*models.Todo); ok {
		r0 = rf(keyword, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(keyword, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *TodoRepository) FindById(id string) (*models.Todo, error) {
	ret := _m.Called(id)

	var r0 *models.Todo
	if rf, ok := ret.Get(0).(func(string) *models.Todo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: value
func (_m *TodoRepository) Store(value *models.Todo) (*models.Todo, error) {
	ret := _m.Called(value)

	var r0 *models.Todo
	if rf, ok := ret.Get(0).(func(*models.Todo) *models.Todo); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Todo) error); ok {
		r1 = rf(value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, value
func (_m *TodoRepository) Update(id string, value *models.Todo) (*models.Todo, error) {
	ret := _m.Called(id, value)

	var r0 *models.Todo
	if rf, ok := ret.Get(0).(func(string, *models.Todo) *models.Todo); ok {
		r0 = rf(id, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *models.Todo) error); ok {
		r1 = rf(id, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
