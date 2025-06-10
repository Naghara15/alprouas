package models

import (
	"alprouas/internal/database"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Id             int       `json:"id" gorm:"PrimaryKey"`
	User_id        int       `json:"user_id"`
	Product_id     int       `json:"product_id"`
	Qty            int       `json:"qty"`
	Total_price    float64   `json:"total_price"`
	Transaction_id *int      `json:"transaction_id"`
	Created_at     time.Time `gorm:"autoCreateTime"`
	Updated_at     time.Time `gorm:"autoUpdateTime"`
}

func AddCart(cart Cart) error {
	var cartExist Cart
	err := database.DB.Where("user_id = ? AND product_id = ? AND transaction_id IS NULL", cart.User_id, cart.Product_id).First(&cartExist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result := database.DB.Create(&cart)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}

	product, err := GetProductByID(cartExist.Product_id)
	if err != nil {
		return err
	}

	cartExist.Total_price += product.Price
	cartExist.Qty += cart.Qty
	result := database.DB.Model(&cartExist).Updates(Cart{Qty: cartExist.Qty, Total_price: cartExist.Total_price})
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func GetCartUser(userID int) ([]Cart, error) {
	var carts []Cart
	result := database.DB.Where("user_id = ? AND transaction_id IS NULL", userID).Order("id").Find(&carts)
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

func GetCartsByID(carts []Cart) ([]Cart, error) {
	ids := make([]int, len(carts))
	for i, s := range carts {
		ids[i] = s.Id
	}

	result := database.DB.Find(&carts, ids)
	if result.Error != nil {
		return []Cart{}, result.Error
	}

	return carts, nil
}

func UpdateCartQty(cart Cart) error {
	product, err := GetProductByID(cart.Product_id)
	if err != nil {
		return err
	}

	cart.Total_price = float64(cart.Qty) * product.Price
	result := database.DB.Model(&cart).Updates(Cart{Qty: cart.Qty, Total_price: cart.Total_price})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func CalculateTotal(carts []Cart) (float64, error) {
	var total float64
	for _, s := range carts {
		cart := Cart{
			Id: s.Id,
		}

		result := database.DB.First(&cart)
		if result.Error != nil {
			return 0, result.Error
		}
		total += cart.Total_price
	}

	return total, nil
}
