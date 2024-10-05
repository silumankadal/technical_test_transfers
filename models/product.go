package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	TransactionType string `json:"transaction_type" form:"transaction_type"`
	Amount          int    `json:"amount" form:"amount" valid:"required~amount of your transaction is required"`
	BalanceBefore   int    `json:"balance_before" form:"balance_before"`
	BalanceAfter    int    `json:"balance_after" form:"balance_after"`
	Remarks         string `json:"remarks" form:"remarks"`
	Payment_id      uint   `json:"payment_id" form:"payment_id"`
	Top_up_id       uint   `json:"top_up_id" form:"top_up_id"`
	Transfer_id     uint   `json:"transfer_id" form:"transfer_id"`
	UserID          uint
	User            *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
