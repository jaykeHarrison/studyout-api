package models

type Users struct {
	//uint can store twice as many *positive* integers as int
	ID       uint   `json:"id" gorm:"primaryKey"`
	username string `json:"username"`
}
