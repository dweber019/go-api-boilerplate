package models

import (
	"errors"
	"time"

	"github.com/dweber019/go-api-boilerplate/app/config"
)

type UserHasUser struct {
	ID        int    		`json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	UserOwns  int    		`json:"userOwns" gorm:"not null"`
	UserLent  int    		`json:"userlent" gorm:"not null"`
	Date 			time.Time `json:"date, omitempty" gorm:"not null"`
	Payed  		bool 			`json:"payed, omitempty"`
	Amount  	float32 	`json:"amount, omitempty"`
}

func (UserHasUser) TableName() string {
	return "user_has_user"
}

func (uu *UserHasUser) FetchAll() []UserHasUser {
	db := config.GetDatabaseConnection()

	var userUsers []UserHasUser
	db.Find(&userUsers)

	return userUsers
}

func (uu *UserHasUser) FetchById() error {
	db := config.GetDatabaseConnection()

	if err := db.Where("id = ?", uu.ID).Find(&uu).Error; err != nil {
		return errors.New("Could not find the users")
	}

	return nil
}

func (uu *UserHasUser) Save() error {
	db := config.GetDatabaseConnection()

	if db.NewRecord(uu) {
		if err := db.Create(&uu).Error; err != nil {
			return errors.New("Could not create users")
		}
	} else {
		if err := db.Save(&uu).Error; err != nil {
			return errors.New("Could not update users")
		}
	}

	return nil
}

func (uu *UserHasUser) Delete() error {
	db := config.GetDatabaseConnection()

	if err := db.Delete(&uu).Error; err != nil {
		return errors.New("Could not find the users")
	}

	return nil
}
