package models

import "alprouas/internal/database"

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
