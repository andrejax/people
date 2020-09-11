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

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}