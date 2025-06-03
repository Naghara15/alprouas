package models

import (
	"alprouas/internal/database"
	"alprouas/internal/utilities"
	"time"
)

type Users struct {
	Id         int       `json:"id" gorm:"PrimaryKey"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
}

func CreateUser(User Users) error {
	var err error
	User.Password, err = utilities.Hash(User.Password)
	if err != nil {
		return err
	}
	result := database.DB.Create(&User)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllUser() ([]Users, error) {
	var User []Users
	result := database.DB.Find(&User)
	if result.Error != nil {
		return nil, result.Error
	}

	return User, nil
}

func GetUserByName(name string) (Users, error) {
	var user Users
	result := database.DB.Where("username = ?", name).First(&user)
	if result.Error != nil {
		return Users{}, result.Error
	}

	return user, nil
}

func GetuserByID(id int) (Users, error) {
	var user Users
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return Users{}, result.Error
	}

	return user, nil
}