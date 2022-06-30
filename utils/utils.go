package utils
import "time"

type Location struct {
	LocationID   uint
	LocationName string
	Address      string
	Postcode     string
	Longitude    float32
	Latitude     float32
	Condition    string
	ImgUrl       string
	CreatedBy    uint
}

type Review struct {
	ReviewID     uint
	UserRefer  uint
	LocationRefer uint
	CreatedAt	time.Time
	VisitDate string
	StarRating int
	ReviewBody string
}