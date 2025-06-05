package models

import (
	"alprouas/internal/database"
)

type Product_tag struct {
	Product_id int `json:"product_id"`
	Tag_id     int `json:"tag_id"`
}

func GetTagsByProductID(product_id int) ([]Product_tag, error) {
	var tags []Product_tag
	result := database.DB.Where("product_id = ?", product_id).Find(&tags)
	if result.Error != nil {
		return []Product_tag{}, result.Error
	}

	return tags, nil
}
