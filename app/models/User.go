package models

import (
	"errors"
	"crypto/md5"
	"encoding/hex"

	"github.com/dweber019/go-api-boilerplate/app/config"
)

type User struct {
	ID        int			`json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Username	string	`json:"username, omitempty" gorm:"not null; type:varchar(100)"`
	Password  string	`json:"password" gorm:"not null; size:512"`
}

func (u *User) SetPassword(p string) {
	hash := md5.Sum([]byte(p))
	u.Password = hex.EncodeToString(hash[:])
}

func (User) TableName() string {
	return "user"
}

func (u *User) FetchAll() []User {
	db := config.GetDatabaseConnection()

	var users []User
	db.Find(&users)

	return users
}

func (u *User) FetchById() error {
	db := config.GetDatabaseConnection()

	if err := db.Where("id = ?", u.ID).Find(&u).Error; err != nil {
		return errors.New("Could not find the user")
	}

	return nil
}

func (u *User) Save() error {
	db := config.GetDatabaseConnection()

	if db.NewRecord(u) {
		if err := db.Create(&u).Error; err != nil {
			return errors.New("Could not create user")
		}
	} else {
		if err := db.Save(&u).Error; err != nil {
			return errors.New("Could not update user")
		}
	}

	return nil
}

func (u *User) Delete() error {
	db := config.GetDatabaseConnection()

	if err := db.Delete(&u).Error; err != nil {
		return errors.New("Could not find the user")
	}

	return nil
}
