package storage

import "go-hello/models"

type MysqlStorage interface {
	GetUserByID(int) *models.User
}
