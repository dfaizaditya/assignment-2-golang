package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `json:"itemCode" valid:"required~Item Code of your item is required"`
	Description string `json:"description" valid:"required~Description of your item is required"`
	Quantity    uint   `json:"quantity" valid:"required~Quantity of your item is required"`
	OrderID     uint
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(i)

	if err != nil {
		err = errCreate
		return
	}

	err = nil

	return
}

func (i *Item) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(i)

	if err != nil {
		err = errCreate
		return
	}

	err = nil

	return
}
