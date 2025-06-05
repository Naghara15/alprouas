package models

import (
	"alprouas/internal/database"
	"errors"

	"gorm.io/gorm"
)

type User_tag struct {
	User_id int `json:"user_id"`
	Tag_id  int `json:"tag_id"`
	Clicked int `json:"clicked"`
}

func GetUserMostClicked(userID int) ([]User_tag, error) {
	var user_tags []User_tag
	result := database.DB.Order("clicked DESC").Limit(3).Find(&user_tags)
	if result.Error != nil {
		return []User_tag{}, nil
	}

	return user_tags, nil
}

func AddClick(user_tag []User_tag) error {
	for _, tag := range user_tag {
		var exist User_tag

		err := database.DB.Where("user_id = ? AND tag_id = ?", tag.User_id, tag.Tag_id).First(&exist).Error
		if err != nil {
			// Buat Jika tidak ada
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tag.Clicked = 1
				if err := database.DB.Create(&tag).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// Jika sudah ada, plus 1 untuk kolom clicked
			tag.Clicked++
			err := database.DB.Model(&User_tag{}).
				Where("user_id = ? AND tag_id = ?", tag.User_id, tag.Tag_id).
				Update("clicked", gorm.Expr("clicked + ?", tag.Clicked)).Error

			if err != nil {
				return err
			}

		}
	}
	return nil
}
