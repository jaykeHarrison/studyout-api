package seed

import (
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/models"
)

func Seed() {
	Users := []models.Users{
		{1, "JaykeHarrison", "Jayke", "Harrison", "jayke@jaykeharrison.com"},
		{2, "GarySum", "Gary", "Sum", "Gary@Gary.com"},
		{3, "JamesYeung", "James", "Yeung", "James@James.com"},
		{4, "CaolanHamilton", "Caolan", "Hamilton", "Caolan@Caolan.com"},
		{5, "NikhilMathew", "Nikhil", "Mathew", "Nikhil@Nikhil.com"},
	}
	Locations := []models.Location{
		{
			LocationName: "19 Cafe Bar",
			Address:      "19 Lever St, Manchester",
			Postcode:     "M1 1BY",
			Longitude:    -2.233770,
			Latitude:     53.482570,
			Condition:    "Buy a coffee at least every 3 hours",
			ImgUrl:       "https://tinyurl.com/42sdhvyd",
			CreatedBy:    1,
		},
		{
			LocationName: "Federal Cafe Bar",
			Address:      "194 Deansgate, Manchester",
			Postcode:     "M3 3ND",
			Longitude:    -2.249450,
			Latitude:     53.478660,
			Condition:    "unlimited stay with any purchase",
			ImgUrl:       "https://tinyurl.com/42sdhvyd",
			CreatedBy:    2,
		},
		{
			LocationName: "Ezra & Gil",
			Address:      "20 Hilton St, Manchester",
			Postcode:     "M1 1FR",
			Longitude:    -2.232888,
			Latitude:     53.482025,
			Condition:    "Can stay upto 4 hours with a purchase",
			ImgUrl:       "https://tinyurl.com/42sdhvyd",
			CreatedBy:    3,
		},
	}

	for _, user := range Users {
		model.AddUser(&user)
	}

	for _, location := range Locations {
		model.AddLocation(&location)
	}
}
