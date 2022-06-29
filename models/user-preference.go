package models

type UserPreference struct {
	UserRefer            int
	Users                Users
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
