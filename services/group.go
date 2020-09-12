package services

import (
	"context"
	"people/interfaces"
	"people/models"
	"time"
)

type GroupService struct {
	groupRepo      interfaces.GroupRepository
	contextTimeout time.Duration
}

func NewGroupService(ur interfaces.GroupRepository, timeout time.Duration) interfaces.GroupService {
	return &GroupService{
		groupRepo:      ur,
		contextTimeout: timeout,
	}
}

func (gr *GroupService) Get(ctx context.Context, id string)  (res models.Group, err error) {
	ctx, cancel := context.WithTimeout(ctx, gr.contextTimeout)
	defer cancel()

	return gr.groupRepo.Get(ctx, id)
}

func (gr *GroupService) List(ctx context.Context)  (res []models.Group, err error) {

	ctx, cancel := context.WithTimeout(ctx, gr.contextTimeout)
	defer cancel()

	return gr.groupRepo.List(ctx)
}

func (gr *GroupService) Add(ctx context.Context, g *models.Group) (err error) {
	ctx, cancel := context.WithTimeout(ctx, gr.contextTimeout)
	defer cancel()
	err = gr.groupRepo.Add(ctx, g)
	return
}

func (gr *GroupService) Update(ctx context.Context, g *models.Group) (err error) {
	ctx, cancel := context.WithTimeout(ctx, gr.contextTimeout)
	defer cancel()

	_, err = gr.Get(ctx, g.Id)
	if err != nil{
		return err
	}

	return gr.groupRepo.Update(ctx, g)
}

func (gr *GroupService) Remove(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, gr.contextTimeout)
	defer cancel()
	_, err = gr.groupRepo.Get(ctx, id)
	if err != nil {
		return
	}

	return gr.groupRepo.Remove(ctx, id)
}