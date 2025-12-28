package repository

import (
	"context"
	"errors"
	"first_service/models"

	"github.com/go-kivik/kivik/v4"
)

type NotebookCounchDB struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rev         string `json:"_rev,omitempty"`
}

type Repository struct {
	couchdb *kivik.DB
}

func NewRepository(c *kivik.DB) *Repository {
	return &Repository{
		couchdb: c,
	}
}

func (r *Repository) Create(ctx context.Context, n *models.NotebookEntityDomain) error {
	NotebookCounch := NotebookCounchDB{
		ID:          n.ID,
		Name:        n.Name,
		Description: n.Description,
		Rev:         "",
	}

	if _, err := r.couchdb.Put(ctx, n.ID, NotebookCounch); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	var nb NotebookCounchDB
	row := r.couchdb.Get(ctx, id)
	if err := row.ScanDoc(&nb); err != nil {
		return err
	}

	rev := nb.Rev // Aqui pega a revisão do documento
	if rev == "" {
		return errors.New("revisão não encontrada.")
	}

	_, err := r.couchdb.Delete(ctx, id, rev)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, id string) (*models.NotebookEntityDomain, error) {
	var nb NotebookCounchDB
	if err := r.couchdb.Get(ctx, id).ScanDoc(&nb); err != nil {
		return nil, err
	}

	nd := &models.NotebookEntityDomain{
		ID:          nb.ID,
		Name:        nb.Name,
		Description: nb.Description,
	}

	return nd, nil
}
