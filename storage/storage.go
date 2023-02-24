package storage

import (
	"go-hello/models"
	"log"
)

type MysqlStorage interface {
	GetUserByID(int) *models.User
}

type InitStore interface {
	InitStorage() error
}

func InitStorages(istore ...InitStore) (err error) {
	IStore := make([]InitStore, 0)
	IStore = append(IStore, istore...)
	for _, store := range IStore {
		err = store.InitStorage()
		if err != nil {
			log.Fatal(err)
		}
	}
	return err
}
