package utils

import (
	"time"

	"github.com/jaykeHarrison/studyout-api/models"
)

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

type Review struct {
	ReviewId      uint
	UserRefer     uint
	LocationRefer uint
	CreatedAt     time.Time
	VisitDate     string
	StarRating    int
	ReviewBody    string
}

func CreateResponseReview(review models.Review) Review {
	return Review{
		ReviewId:      review.ReviewId,
		UserRefer:     review.UserRefer,
		LocationRefer: review.LocationRefer,
		CreatedAt:     review.CreatedAt,
		VisitDate:     review.VisitDate,
		StarRating:    review.StarRating,
		ReviewBody:    review.ReviewBody,
	}
}

type Location struct {
	LocationID   uint    `gorm:"primaryKey"`
	LocationName string  `gorm:"not null; default:null"`
	Address      string  `gorm:"not null; default:null"`
	Postcode     string  `gorm:"not null; default:null"`
	Longitude    float32 `gorm:"not null; default:null"`
	Latitude     float32 `gorm:"not null; default:null"`
	Condition    string  `gorm:"not null; default:null"`
	ImgUrl       string  `gorm:"not null; default:null"`
	CreatedBy    uint    `gorm:"not null; default:null"`
	Users        User    `gorm:"foreignKey:CreatedBy"`
}

func CreateResponseLocation(location models.Location) Location {
	return Location{LocationID: location.LocationID, LocationName: location.LocationName, Address: location.Address, Postcode: location.Postcode, Longitude: location.Longitude, Latitude: location.Latitude, Condition: location.Condition, ImgUrl: location.ImgUrl, CreatedBy: location.CreatedBy}
}
