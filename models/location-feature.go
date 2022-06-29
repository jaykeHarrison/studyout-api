package models

type LocationFeature struct {
	LocationId           uint
	Location             Location `gorm:"foreignKey:LocationId"`
	WiFi                 bool
	WaterFountain        bool
	PowerSockets         bool
	Noisy                bool
	Food                 bool
	Beverages            bool
	Toilets              bool
	WheelchairAccess     bool
	DiasabledToilets     bool
	GenderNeutralToilets bool
	FreeParking          bool
	PaidParking          bool
	PetFriendly          bool
	Vegetarian           bool
	Vegan                bool
	OutdoorSeating       bool
	PhoneSignal          bool
	NaturalLight         bool
}
