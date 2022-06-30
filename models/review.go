package models

import "time"

type Review struct {
	UserRefer     uint `gorm:"not null; default:null"`
	LocationRefer uint `gorm:"not null; default:null"`
	ReviewId      uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UserId        Users    `gorm:"foreignKey:UserRefer; not null; default:null"`
	LocationId    Location `gorm:"foreignKey:LocationRefer; not null; default:null"`
	VisitDate     string   `gorm:"not null; default:null"`
	StarRating    int      `gorm:"not null; default:null"`
	// Another ticket to handle the range of the star rating
	ReviewBody string `gorm:"not null; default:null"`
}
