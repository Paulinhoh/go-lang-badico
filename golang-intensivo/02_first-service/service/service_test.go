package service_test

import (
	"context"
	"first_service/models"
	"first_service/repository"
	"first_service/service"
	"log"
	"testing"

	"github.com/go-kivik/kivik/v4"
	"github.com/stretchr/testify/require"
	"github.com/tj/assert"
)

func build() *service.Service {
	client, err := kivik.New("couch", "http://admin:pass@localhost:5984/")
	if err != nil {
		log.Fatalf("erro ao criar cliente CouchDB: %s", err)
	}

	db := client.DB("notebooks")
	if err := db.Err(); err != nil {
		log.Fatalf("erro ao conectar ao DB: %s", err)
	}

	repo := repository.NewRepository(db)
	srv := service.NewService(repo)

	return srv
}

func TestService(t *testing.T) {
	srv := build()
	if srv == nil {
		t.Error("srv tem que existir!")
	}

	t.Run("Service Create", func(t *testing.T) {
		// input
		input := models.CreateNotebookInputDTO{
			Name:        "Test Notebook",
			Description: "Notebook de teste",
		}

		// input srv.Create
		ctx := context.TODO()
		output, err := srv.Create(ctx, input)
		require.NoError(t, err)

		// output
		assert.Equal(t, input.Name, output.Name)
		assert.Equal(t, input.Description, output.Description)
		assert.NotEmpty(t, output.ID)
	})
}
