package models

type Users struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null; default:null"`
	FirstName string `gorm:"not null; default:null"`
	LastName  string `gorm:"not null; default:null"`
	Email     string `gorm:"not null; default:null"`
}
