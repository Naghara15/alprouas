package models

import (
	"alprouas/internal/database"
	"time"
)

type Product struct {
	Id          int       `json:"id" gorm:"PrimaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price 		float64   `json:"price"`
	Stock       int       `json:"stock"`
	Image		string	  `json:"image"`
	Rating		float64	  `json:"rating"`
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at  time.Time `gorm:"autoUpdateTime"`
}

func SearchProducts(name string) ([]Product, error) {
	var products []Product
	result := database.DB.Where("name LIKE ?", "%"+name+"%").Find(&products)
	if result.Error != nil {
		return []Product{}, result.Error
	}

	return products, nil
}

func GetAllProducts() ([]Product, error) {
	var products []Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		return []Product{}, result.Error
	}

	return products, nil
}

func GetProductByID(id int) (Product, error) {
	var product Product
	result := database.DB.Find(&product, id)
	if result.Error != nil {
		return Product{}, result.Error
	}

	return product, nil
}

func GetProductsByTag(tag []int) ([]Product, error) {
	var product_tag []Product_tag

	findRel := database.DB.Where("tag_id IN ?", tag).Find(&product_tag)
	if findRel.Error != nil {
		return []Product{}, findRel.Error
	}

	var productIDs []int
	for _, s := range product_tag {
		productIDs = append(productIDs, s.Product_id)
	}

	var products []Product
	result := database.DB.Where("id IN ?", productIDs).Find(&products)
	if result.Error != nil {
		return []Product{}, nil
	}

	return products, nil
}

func CreateProduct(product Product) error {
	result := database.DB.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UseStock(carts []Cart) error {
	for _, s := range carts {
		var product Product
		availProduct := database.DB.First(&product, s.Product_id)
		if availProduct.Error != nil {
			return availProduct.Error
		}

		product.Stock -= s.Qty
		result := database.DB.Save(&product)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
