package models

import (
	"alprouas/internal/database"
	"time"
)

type Cart struct {
	Id          	int       	`json:"id" gorm:"PrimaryKey"`
	User_id     	int       	`json:"user_id"`
	Product_id  	int       	`json:"product_id"`
	Qty         	int       	`json:"qty"`
	Total_price 	float64   	`json:"total_price"`
	Transaction_id *int	  		`json:"transaction_id"`
	Created_at  	time.Time 	`gorm:"autoCreateTime"`
	Updated_at  	time.Time 	`gorm:"autoUpdateTime"`
}

func AddCart(cart Cart) error {
	result := database.DB.Create(&cart)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetCartUser(userID int) ([]Cart, error) {
	var carts []Cart
	result := database.DB.Where("user_id = ? AND transaction_id IS NULL", userID).Find(&carts)
	if result.Error != nil {
		return []Cart{}, result.Error
	}

	return carts, nil
}

func DeleteCart(id int) error {
	result := database.DB.Delete(&Cart{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateTrID(carts []Cart) error {
	for _, cart := range carts {
		err := database.DB.Model(&Cart{}).
			Where("id = ?", cart.Id).
			Update("transaction_id", cart.Transaction_id).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func GetCartsByID(carts []Cart) ([]Cart,error) {
	ids := make([]int, len(carts))
	for i,s := range carts{
		ids[i] = s.Id
	}

	result := database.DB.Find(&carts, ids)
	if result.Error != nil {
		return []Cart{}, result.Error
	}

	return carts, nil
}