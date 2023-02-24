package storage

import (
	"fmt"
	"go-hello/conf"
	"go-hello/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlStore struct {
	mysql_user     string
	mysql_password string
	mysql_Host     string
	mysql_port     string
	db_Name        string
	db             *gorm.DB
}

func NewMysqlStorage() *MySqlStore {
	mysql_conf := conf.Conf.MysqlConfig
	return &MySqlStore{
		mysql_user:     mysql_conf.Mysql_User,
		mysql_password: mysql_conf.Password,
		mysql_Host:     mysql_conf.Mysql_Host,
		mysql_port:     mysql_conf.Mysql_port,
		db_Name:        mysql_conf.DB_Name,
		db:             &gorm.DB{},
	}
}

type Intsence struct {
	MSotre *MySqlStore
}

func (i Intsence) InitStorage() error {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", i.MSotre.mysql_user, i.MSotre.mysql_password, i.MSotre.mysql_Host, i.MSotre.mysql_port, i.MSotre.db_Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	i.MSotre.db = db
	return err
}
func (s *MySqlStore) GetUserByID(identity int) *models.User {
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
