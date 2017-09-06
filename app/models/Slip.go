package models

import (
	"errors"
	"time"

	"github.com/dweber019/go-api-boilerplate/app/config"
)

type Slip struct {
	ID        int    		`json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Date 			time.Time `json:"date, omitempty" gorm:"not null"`
	Price  		float32 	`json:"price, omitempty"`
	Payed  		bool 			`json:"payed, omitempty"`
	User  		User 			`json:"user, omitempty" gorm:"ForeignKey:UserID;AssociationForeignKey:Refer"`
	UserID  	int 			`json:"userID, omitempty"`
}

func (Slip) TableName() string {
	return "slip"
}

func (s *Slip) FetchAll() []Slip {
	db := config.GetDatabaseConnection()

	var slips []Slip
	db.Find(&slips)

	return slips
}

func (s *Slip) FetchById() error {
	db := config.GetDatabaseConnection()

	if err := db.Where("id = ?", s.ID).Find(&s).Error; err != nil {
		return errors.New("Could not find the slip")
	}

	return nil
}

func (s *Slip) Save() error {
	db := config.GetDatabaseConnection()

	if db.NewRecord(s) {
		if err := db.Create(&s).Error; err != nil {
			return errors.New("Could not create slip")
		}
	} else {
		if err := db.Save(&s).Error; err != nil {
			return errors.New("Could not update slip")
		}
	}

	return nil
}

func (s *Slip) Delete() error {
	db := config.GetDatabaseConnection()

	if err := db.Delete(&s).Error; err != nil {
		return errors.New("Could not find the slip")
	}

	return nil
}
