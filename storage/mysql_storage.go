package storage

import (
	"fmt"
	"go-hello/conf"
	"go-hello/models"
	"go-hello/utils"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dB *gorm.DB

type MySqlStore struct {
	mysql_user     string
	mysql_password string
	mysql_Host     string
	mysql_port     string
	db_Name        string
}

func NewMysqlStorage() *MySqlStore {
	mysql_conf := conf.Conf.MysqlConfig
	return &MySqlStore{
		mysql_user:     mysql_conf.Mysql_User,
		mysql_password: mysql_conf.Password,
		mysql_Host:     mysql_conf.Mysql_Host,
		mysql_port:     mysql_conf.Mysql_port,
		db_Name:        mysql_conf.DB_Name,
	}
}

type Intsence struct {
	MSotre *MySqlStore
}

func (i Intsence) InitStorage() error {
	i.MSotre = NewMysqlStorage()
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", i.MSotre.mysql_user, i.MSotre.mysql_password, i.MSotre.mysql_Host, i.MSotre.mysql_port, i.MSotre.db_Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	dB = db
	dB.AutoMigrate(&models.User{})
	return err
}
func (s *MySqlStore) GetUserByID(identity int64) *models.User {
	return &models.User{
		Username: "bob",
		Password: "haha",
		Email:    "122222",
		Content:  "halo how are u",
		Birthday: "2020/1/11",
		Company:  "CJNJ",
		UserJob:  "drug deler",
		City:     "juadalahala",
		Number:   12223333333,
	}
}

func (s *MySqlStore) UserRegister(userRegister *models.UserRegister) error {
	user := &models.User{
		Username: userRegister.Username,
		Password: utils.EncryptPassword(userRegister.Password),
		Email:    userRegister.Emial,
	}
	for {
		number := utils.GenUserNumber()
		result := dB.Where("number=?", number).Find(user)
		if result.RowsAffected == 0 {
			user.Number = number
			break
		}
	}
	restult := dB.Create(user)
	return restult.Error
}

func (s *MySqlStore) GetUserByNameAndEmail(name string, email string) (int64, error) {
	user := &models.User{}
	result := dB.Where("username=? or email=?", name, email).Find(user)
	return result.RowsAffected, result.Error
}
