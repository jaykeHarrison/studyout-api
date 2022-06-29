package utils

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
