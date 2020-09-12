package tests

import (
	"context"
	"github.com/stretchr/testify/mock"
	"people/models"
)

type UserService struct {
	mock.Mock
}

func (_m *UserService) Add(_a0 context.Context, _a1 *models.User) (err error) {
	ret := _m.Called(_a0, _a1)
	return ret.Error(0)
}

func (_m *UserService) Remove(_a0 context.Context, _a1 string) (err error) {
	ret := _m.Called(_a0, _a1)
	return ret.Error(0)
}

func (_m *UserService) Get(_a0 context.Context, _a1 string) (res models.User, err error) {
	ret := _m.Called(_a0, _a1)
	res = ret.Get(0).(models.User)
	err = ret.Error(1)
	return
}

func (_m *UserService) Update(_a0 context.Context, _a1 *models.User) (err error) {
	ret := _m.Called(_a0, _a1)
	return ret.Error(0)
}

func (_m *UserService) List(_a0 context.Context) (res []models.User, err error)  {
	ret := _m.Called(_a0)
	res = ret.Get(0).([]models.User)
	err = ret.Error(1)
	return
}