package storage

import (
	"go-hello/models"
	"time"
)

type MySqlStore struct{}

func NewMysqlStorage() *MySqlStore {
	return &MySqlStore{}
}
func (s MySqlStore) GetUserByID(identity int) *models.User {
	return &models.User{
		Username: "bob",
		Password: "haha",
		Email:    "122222",
		Content:  "halo how are u",
		Birthday: time.Now(),
		Company:  "CJNJ",
		UserJob:  "drug deler",
		City:     "juadalahala",
		Identity: identity,
	}
}
