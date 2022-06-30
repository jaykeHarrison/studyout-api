package models

type Location struct {
    LocationID   uint  `gorm:"primaryKey"`
    LocationName string  `gorm:"not null; default:null"`
    Address      string  `gorm:"not null; default:null"`
    Postcode     string  `gorm:"not null; default:null"`
    Longitude    float32  `gorm:"not null; default:null"`
    Latitude     float32  `gorm:"not null; default:null"`
    Condition    string  `gorm:"not null; default:null"`
    ImgUrl       string  `gorm:"not null; default:null"`
    CreatedBy    uint   `gorm:"not null; default:null"`
    Users        Users `gorm:"foreignKey:CreatedBy"`
}