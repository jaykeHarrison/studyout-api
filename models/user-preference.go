package models

type UserPreference struct {
	UserRefer            int   `gorm: not null; default:null`
	Users                Users `gorm:"foreignKey:UserRefer"`
	WiFi                 bool  `gorm: type:bool; default:false`
	WaterFountain        bool  `gorm: type:bool; default:false`
	PowerSockets         bool  `gorm: type:bool; default:false`
	Noisy                bool  `gorm: type:bool; default:false`
	Food                 bool  `gorm: type:bool; default:false`
	Beverages            bool  `gorm: type:bool; default:false`
	Toilets              bool  `gorm: type:bool; default:false`
	WheelchairAccess     bool  `gorm: type:bool; default:false`
	DiasabledToilets     bool  `gorm: type:bool; default:false`
	GenderNeutralToilets bool  `gorm: type:bool; default:false`
	FreeParking          bool  `gorm: type:bool; default:false`
	PaidParking          bool  `gorm: type:bool; default:false`
	PetFriendly          bool  `gorm: type:bool; default:false`
	Vegetarian           bool  `gorm: type:bool; default:false`
	Vegan                bool  `gorm: type:bool; default:false`
	OutdoorSeating       bool  `gorm: type:bool; default:false`
	PhoneSignal          bool  `gorm: type:bool; default:false`
	NaturalLight         bool  `gorm: type:bool; default:false`
}
