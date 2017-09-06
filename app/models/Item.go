package models

import (
	"errors"

	"github.com/dweber019/go-api-boilerplate/app/config"
)

type Item struct {
	ID        int    		`json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Name 			string 		`json:"name, omitempty" gorm:"not null"`
}

func (Item) TableName() string {
	return "item"
}

func (i *Item) FetchAll() []Item {
	db := config.GetDatabaseConnection()

	var items []Item
	db.Find(&items)

	return items
}

func (i *Item) FetchById() error {
	db := config.GetDatabaseConnection()

	if err := db.Where("id = ?", i.ID).Find(&i).Error; err != nil {
		return errors.New("Could not find the item")
	}

	return nil
}

func (i *Item) Save() error {
	db := config.GetDatabaseConnection()

	if db.NewRecord(i) {
		if err := db.Create(&i).Error; err != nil {
			return errors.New("Could not create item")
		}
	} else {
		if err := db.Save(&i).Error; err != nil {
			return errors.New("Could not update item")
		}
	}

	return nil
}

func (i *Item) Delete() error {
	db := config.GetDatabaseConnection()

	if err := db.Delete(&i).Error; err != nil {
		return errors.New("Could not find the item")
	}

	return nil
}
