package utils

import "github.com/jaykeHarrison/studyout-api/models"

type Location struct {
	LocationID   uint
	LocationName string
	Address      string
	Postcode     string
	Longitude    float32
	Latitude     float32
	Condition    string
	ImgUrl       string
	CreatedBy    uint
}

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	FirstName string
	LastName  string
	Email     string
}

func CreateResponseUser(user models.Users) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}
