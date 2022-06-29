package models

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
	Users        Users `gorm:"foreignKey:CreatedBy"`
}
