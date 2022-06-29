package models

import "time"

type Review struct {
	UserRefer  uint
	LocationRefer uint
	ReviewId	uint	`gorm:"primaryKey"`
	CreatedAt	time.Time
	UserId  Users `gorm:"foreignKey:UserRefer"`
	LocationId	Location `gorm:"foreignKey:LocationRefer"`
	VisitDate string
	StarRating int
	// Another ticket to handle the range of the star rating
	ReviewBody string
}