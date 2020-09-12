package interfaces

import (
	"context"
)
import "people/models"

type UserRepository interface {
	Get(ctx context.Context, id string) (models.User, error)
	List(ctx context.Context) ([]models.User, error)
	Add(ctx context.Context, a *models.User) error
	Update(ctx context.Context, u *models.User) error
	Remove(ctx context.Context, id string) error
}