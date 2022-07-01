package models

import "time"

type Review struct {
	UserRefer     uint `gorm:"not null; default:null"`
	LocationRefer uint `gorm:"not null; default:null"`
	ReviewId      uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UserId        Users    `gorm:"foreignKey:UserRefer"`
	LocationId    Location `gorm:"foreignKey:LocationRefer"`
	VisitDate     string   `gorm:"not null; default:null"`
	StarRating    int      `gorm:"not null; default:null"`
	ReviewBody    string   `gorm:"not null; default:null"`
}
