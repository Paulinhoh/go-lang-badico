package service

import (
	"context"
	"first_service/models"
	"first_service/repository"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		repo: r,
	}
}

// Create
func (srv *Service) Create(ctx context.Context, input models.CreateNotebookInputDTO) (*models.NotebookEntityDomain, error) {
	id := uuid.NewString()

	newNotebook := &models.NotebookEntityDomain{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}

	if err := srv.repo.Create(ctx, newNotebook); err != nil {
		return nil, err
	}

	return newNotebook, nil
}

// Delete
func (srv *Service) Delete(ctx context.Context, id string) error {
	if err := srv.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

// Get
func (srv *Service) Get(ctx context.Context, id string) (*models.NotebookEntityDomain, error) {
	nb, err := srv.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return nb, nil
}
