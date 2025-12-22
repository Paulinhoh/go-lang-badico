package service

import (
	"errors"

	"github.com/Paulinhoh/golang-basico/aula_04/10_api-project/model"
)

var users = []model.User{
	{ID: "1", Name: "Alice", Age: 25},
	{ID: "2", Name: "Bob", Age: 30},
}

func GetAllUsers() []model.User {
	return users
}

func GetUserByID(id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("usuario não encontrado")
}

// service -> usa os models
