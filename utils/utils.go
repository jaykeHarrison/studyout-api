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
	Features     LocationFeature
}

type LocationFeature struct {
	WiFi                 bool `gorm: type:bool; default:false`
	WaterFountain        bool `gorm: type:bool; default:false`
	PowerSockets         bool `gorm: type:bool; default:false`
	Noisy                bool `gorm: type:bool; default:false`
	Food                 bool `gorm: type:bool; default:false`
	Beverages            bool `gorm: type:bool; default:false`
	Toilets              bool `gorm: type:bool; default:false`
	WheelchairAccess     bool `gorm: type:bool; default:false`
	DiasabledToilets     bool `gorm: type:bool; default:false`
	GenderNeutralToilets bool `gorm: type:bool; default:false`
	FreeParking          bool `gorm: type:bool; default:false`
	PaidParking          bool `gorm: type:bool; default:false`
	PetFriendly          bool `gorm: type:bool; default:false`
	Vegetarian           bool `gorm: type:bool; default:false`
	Vegan                bool `gorm: type:bool; default:false`
	OutdoorSeating       bool `gorm: type:bool; default:false`
	PhoneSignal          bool `gorm: type:bool; default:false`
	NaturalLight         bool `gorm: type:bool; default:false`
}

func CreateResponseLocationFeature(locationFeature models.LocationFeature) LocationFeature {
	return LocationFeature{
		WiFi:                 locationFeature.WiFi,
		WaterFountain:        locationFeature.WaterFountain,
		PowerSockets:         locationFeature.PowerSockets,
		Noisy:                locationFeature.Noisy,
		Food:                 locationFeature.Food,
		Beverages:            locationFeature.Beverages,
		Toilets:              locationFeature.Toilets,
		WheelchairAccess:     locationFeature.WheelchairAccess,
		DiasabledToilets:     locationFeature.DiasabledToilets,
		GenderNeutralToilets: locationFeature.GenderNeutralToilets,
		FreeParking:          locationFeature.FreeParking,
		PaidParking:          locationFeature.PaidParking,
		PetFriendly:          locationFeature.PetFriendly,
		Vegetarian:           locationFeature.Vegetarian,
		Vegan:                locationFeature.Vegan,
		OutdoorSeating:       locationFeature.OutdoorSeating,
		PhoneSignal:          locationFeature.PhoneSignal,
		NaturalLight:         locationFeature.NaturalLight,
	}
}

func CreateResponseLocation(location models.Location, locationFeature models.LocationFeature) LocationResponse {
	responseLocationFeature := CreateResponseLocationFeature(locationFeature)

	return LocationResponse{
		LocationID:   location.LocationID,
		LocationName: location.LocationName,
		Address:      location.Address,
		Postcode:     location.Postcode,
		Longitude:    location.Longitude,
		Latitude:     location.Latitude,
		Condition:    location.Condition,
		ImgUrl:       location.ImgUrl,
		CreatedBy:    location.Users.Username,
		Features:     responseLocationFeature,
	}
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
