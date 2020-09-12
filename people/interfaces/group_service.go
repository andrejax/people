package interfaces

import (
	"context"
	"people/models"
)

type GroupService interface
{
	Get(ctx context.Context, id string) (models.Group, error)
	List(ctx context.Context) ([]models.Group, error)
	Add(ctx context.Context, a *models.Group) error
	Update(ctx context.Context, u *models.Group) error
	Remove(ctx context.Context, id string) error
}