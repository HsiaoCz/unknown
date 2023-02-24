package storage

import (
	"go-hello/models"
	"log"
)

type MysqlStorage interface {
	GetUserByNameAndEmail(string, string) (int64, error)
	UserRegister(*models.UserRegister) error
	GetUserByID(int64) *models.User
	UserSignup(string, string) int64
}

// init storage interface
// use it to connect the storage
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
