package models

import (
	"errors"

	"github.com/dweber019/go-api-boilerplate/app/config"
)

type SlipHasItem struct {
	SlipID        int    `json:"slipID" gorm:"not null"`
	ItemID        int    `json:"itemID" gorm:"not null"`
	Description		string `json:"description, omitempty" gorm:"size:512"`
}

func (SlipHasItem) TableName() string {
	return "slip_has_item"
}

func (si *SlipHasItem) FetchAll() []SlipHasItem {
	db := config.GetDatabaseConnection()

	var slipItems []SlipHasItem
	db.Find(&slipItems)

	return slipItems
}

func (si *SlipHasItem) FetchBySlipId() error {
	db := config.GetDatabaseConnection()

	if err := db.Where("slip_id = ?", si.SlipID).Find(&si).Error; err != nil {
		return errors.New("Could not find the slip item")
	}

	return nil
}

func (si *SlipHasItem) Save() error {
	db := config.GetDatabaseConnection()

	if db.NewRecord(si) {
		if err := db.Create(&si).Error; err != nil {
			return errors.New("Could not create slip item")
		}
	} else {
		if err := db.Save(&si).Error; err != nil {
			return errors.New("Could not update slip item")
		}
	}

	return nil
}

func (si *SlipHasItem) Delete() error {
	db := config.GetDatabaseConnection()

	if err := db.Delete(&si).Error; err != nil {
		return errors.New("Could not find the slip item")
	}

	return nil
}
