package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"not null" json:"customerName" valid:"required~Your customer name is required"`
	OrderedAt    time.Time `json:"orderedAt,omitempty"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(o)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil

	return
}

func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(o)

	if err != nil {
		err = errCreate
		return
	}

	err = nil

	return
}
