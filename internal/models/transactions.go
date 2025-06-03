package models

import (
	"alprouas/internal/database"
	"time"
)

type Transaction struct {
	Id          int       `json:"id" gorm:"PrimaryKey"`
	User_id     int       `json:"user_id"`
	Product_id  int       `json:"product_id"`
	Qty         int       `json:"qty"`
	Total_price float64   `json:"total_price"`
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at  time.Time `gorm:"autoUpdateTime"`
}

func Buy(carts []Cart) error {
	transactions := make([]Transaction, len(carts))
	var err error
	for i, s := range carts {
		transactions[i].User_id = s.User_id
		transactions[i].Product_id = s.Product_id
		transactions[i].Qty = s.Qty
		transactions[i].Total_price = s.Total_price
	}

	result := database.DB.Create(&transactions)
	if result.Error != nil {
		return result.Error
	}

	err = UseStock(carts)
	if err != nil {
		return err
	}

	for i := range carts {
		carts[i].Transaction_id = &transactions[i].Id
	}
	err = UpdateTrID(carts)
	if err != nil {
		return err
	}

	return nil
}
