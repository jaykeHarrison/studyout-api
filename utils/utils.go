package utils

import (
	"time"

	"github.com/jaykeHarrison/studyout-api/models"
)

type User struct {
	ID        uint
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
	LocationID   uint `gorm:"primaryKey"`
	LocationName string
	Address      string
	Postcode     string
	Longitude    float32
	Latitude     float32
	Condition    string
	ImgUrl       string
	CreatedBy    uint
	Users        User
}

type LocationResponse struct {
	LocationID   uint
	LocationName string
	Address      string
	Postcode     string
	Longitude    float32
	Latitude     float32
	Condition    string
	ImgUrl       string
	CreatedBy    string
}

func CreateResponseLocation(location models.Location) LocationResponse {
	return LocationResponse{LocationID: location.LocationID, LocationName: location.LocationName, Address: location.Address, Postcode: location.Postcode, Longitude: location.Longitude, Latitude: location.Latitude, Condition: location.Condition, ImgUrl: location.ImgUrl, CreatedBy: location.Users.Username}
}

type BookmarkResponse struct {
	UserId     uint
	LocationId uint
}

type Bookmark struct {
	UserId     uint64
	LocationId uint64
}

func CreateResponseBookmark(bookmark models.Bookmark) BookmarkResponse {
	return BookmarkResponse{LocationId: bookmark.LocationId, UserId: bookmark.UserId}
}

type LocationPost struct {
	Location        models.Location
	LocationFeature models.LocationFeature
}
