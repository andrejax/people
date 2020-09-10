package services

import (
	"context"
	"people/interfaces"
	"people/models"
	"time"
)

type UserService struct {
	userRepo interfaces.UserRepository
	contextTimeout time.Duration
}

func NewUserService(ur interfaces.UserRepository, timeout time.Duration) interfaces.UserService {
	return &UserService{
		userRepo: ur,
		contextTimeout: timeout,
	}
}

func (us *UserService) Get(ctx context.Context, id int64)  (res models.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	return us.userRepo.Get(ctx, id)
}

func (us *UserService) List(ctx context.Context)  (res []models.User, err error) {

	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()

	return us.userRepo.List(ctx)
}

func (us *UserService) Add(ctx context.Context, u *models.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	err = us.userRepo.Add(ctx, u)
	return
}

func (us *UserService) Update(ctx context.Context, u *models.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	return us.userRepo.Update(ctx, u)
}

func (us *UserService) Remove(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	_, err = us.userRepo.Get(ctx, id)
	if err != nil {
		return
	}

	return us.userRepo.Remove(ctx, id)
}