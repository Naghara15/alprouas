package models

import (
	"alprouas/internal/database"
	"time"
)

type Tag struct {
	Id int `json:"id" gorm:"PrimaryKey"`
	Name string `json:"name"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
}

func GetAllTags() ([]Tag, error){
	var tags []Tag
	result := database.DB.Find(&tags)
	if result.Error != nil {
		return []Tag{}, result.Error
	}

	return tags, nil
}

