package models

type Users struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	FirstName string
	LastName  string
	Email     string
}
