package interfaces

import (
	"context"
	"people/models"
)

type UserService interface
{
	Get(ctx context.Context, id int64) (models.User, error)
	List(ctx context.Context) ([]models.User, error)
	Add(ctx context.Context, a *models.User) error
	Update(ctx context.Context, u *models.User) error
	Remove(ctx context.Context, id int64) error
}