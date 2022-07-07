package seed

import (
	"github.com/jaykeHarrison/studyout-api/model"
	"github.com/jaykeHarrison/studyout-api/models"
)

func Seed() {
	Users := []models.Users{
		{101, "JaykeHarrison", "Jayke", "Harrison", "jayke@jaykeharrison.com"},
		{102, "GarySum", "Gary", "Sum", "Gary@Gary.com"},
		{103, "JamesYeung", "James", "Yeung", "James@James.com"},
		{104, "CaolanHamilton", "Caolan", "Hamilton", "Caolan@Caolan.com"},
		{105, "NikhilMathew", "Nikhil", "Mathew", "Nikhil@Nikhil.com"},
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
			CreatedBy:    101,
		},
		{
			LocationName: "Federal Cafe Bar",
			Address:      "194 Deansgate, Manchester",
			Postcode:     "M3 3ND",
			Longitude:    -2.249450,
			Latitude:     53.478660,
			Condition:    "unlimited stay with any purchase",
			ImgUrl:       "https://tinyurl.com/42sdhvyd",
			CreatedBy:    102,
		},
		{
			LocationName: "Ezra & Gil",
			Address:      "20 Hilton St, Manchester",
			Postcode:     "M1 1FR",
			Longitude:    -2.232888,
			Latitude:     53.482025,
			Condition:    "Can stay upto 4 hours with a purchase",
			ImgUrl:       "https://tinyurl.com/42sdhvyd",
			CreatedBy:    105,
		},
		{
			LocationName: "BrewDog",
			Address:      "144 Oxford Rd, Manchester",
			Postcode:     "M13 9GP",
			Longitude:    -2.237423,
			Latitude:     53.4678349,
			Condition:    "£10 for unlimited tea/coffee/water and a pint. Stay from open to 5pm",
			ImgUrl:       "https://brewdog.mtchost.com/public/files/BLOG%20PHOTOS/OutpostManchester_4.jpg",
			CreatedBy:    104,
		},
		{
			LocationName: "Ground Floor @ Mercure",
			Address:      "Portland St, Manchester ",
			Postcode:     "M1 4PH",
			Longitude:    -2.237360617,
			Latitude:     53.47978283,
			Condition:    "Few desks and tables on G/F of the hotel. 24 hours 7 days free access. No purchase is required. Stay as long as you can.",
			ImgUrl:       "https://cf.bstatic.com/xdata/images/hotel/max300/313033278.jpg?k=9ba0e5d795be9866e42a9232564148fadef5ac928e75b9c554fa9395ff73b758&o=&hp=1",
			CreatedBy:    103,
		},
		{
			LocationName: "The Koffee Pot",
			Address:      "84-86 Oldham St, Manchester",
			Postcode:     "M4 1LE",
			Longitude:    -2.2507945,
			Latitude:     53.4841498,
			Condition:    "Upto 3 hours stay with any food or drink purchase",
			ImgUrl:       "https://lh5.googleusercontent.com/p/AF1QipPG8fVCl4B1X8ycu5MKFSavFw_JyJLphEKKGuzr=w426-h240-k-no",
			CreatedBy:    102,
		},
		{
			LocationName: "Reflexion Lounge",
			Address:      "252 The Quays, Salford",
			Postcode:     "M50 3SD",
			Longitude:    -2.3053677,
			Latitude:     53.4741704,
			Condition:    "No restriction on outdoor seating but 2 hour stay indoor with any purchase",
			ImgUrl:       "https://lh5.googleusercontent.com/p/AF1QipM1fCPSG0d_P-hOn0rns4L15kUG_4WjYZZ0ZEvd=w480-h240-k-no",
			CreatedBy:    101,
		},
		{
			LocationName: "The Ford Madox Brown",
			Address:      "Wilmslow Park, Wilmslow Rd, Manchester",
			Postcode:     "M13 9NG",
			Longitude:    -2.248842,
			Latitude:     53.4658658,
			Condition:    "Stay as long as you want, no purchase required",
			ImgUrl:       "https://lh5.googleusercontent.com/p/AF1QipNfhRSqhSDDz3iMjMXVWSivd1-0tbCZIwLDhQHP=w427-h240-k-no",
			CreatedBy:    105,
		},
		{
			LocationName: "The Whitworth Café",
			Address:      "The Whitworth The University of Manchester, Oxford Rd, Manchester M15 6ER",
			Postcode:     "M15 6ER",
			Longitude:    -2.229871215,
			Latitude:     53.46013361,
			Condition:    "Offering warm, friendly hospitality, together with fresh local and seasonal food.",
			ImgUrl:       "https://www.hospitality.manchester.ac.uk/thewhitworthcafe/CafePicture_resized.jpg",
			CreatedBy:    104,
		},
		{
			LocationName: "Terry's Cafe London",
			Address:      "158 Great Suffolk St, London",
			Postcode:     "SE1 1PE",
			Longitude:    -0.1680209,
			Latitude:     51.5080499,
			Condition:    "Owners very accommodating, can stay as long as you want with a purchase",
			ImgUrl:       "https://lh5.googleusercontent.com/p/AF1QipN4Wkye5A_dHep72CQm2MvgT6nsWp61flQA5qqF=w408-h408-k-no",
			CreatedBy:    103,
		},
		{
			LocationName: "The Monocle Café",
			Address:      "18 Chiltern St, London",
			Postcode:     "W1U 7QA",
			Longitude:    -0.2246437,
			Latitude:     51.5187934,
			Condition:    "Very small but lovely cafe. No more than an hour stay with purchase when busy",
			ImgUrl:       "https://lh5.googleusercontent.com/p/AF1QipOuvDHuVCL9f4A_ECrb711Db1qEFGf-C-6-jrGd=w408-h306-k-no",
			CreatedBy:    102,
		},
		{
			LocationName: "Central Library",
			Address:      "St Peter's Square, Manchester",
			Postcode:     "M2 5PD",
			Longitude:    -2.244105678,
			Latitude:     53.47835695,
			Condition:    "This is quite possibly the fanciest place to study in Manchester. No purchase is required but there is a cafe to get some refreshments",
			ImgUrl:       "https://oxfordroadcorridor.com/wp-content/uploads/2021/10/e743ae886792a296e6b36d175bd1f4ab.jpg",
			CreatedBy:    101,
		},
		{
			LocationName: "Manchester Technology Centre",
			Address:      "103 Oxford Rd, Manchester",
			Postcode:     "M1 7ED",
			Longitude:    -2.238327215,
			Latitude:     53.47236293,
			Condition:    "A lot of space you can use in the public area on ground floor. You can stay as long as you want.",
			ImgUrl:       "https://www.prolificnorth.co.uk/sites/default/files/styles/hero/public/images/features/msptechincubatormay180272.jpg?itok=EKfLzLwe",
			CreatedBy:    105,
		},
		{
			LocationName: "Hideaway Coffee",
			Address:      "7 Farrier's Psge, London",
			Postcode:     "W1D 7DP",
			Longitude:    -0.1351975577,
			Latitude:     51.51144205,
			Condition:    "Make a purchase every 2 hours",
			ImgUrl:       "https://media-cdn.tripadvisor.com/media/photo-s/15/6f/ff/af/the-facade.jpg",
			CreatedBy:    104,
		},
		{
			LocationName: "WatchHouse",
			Address:      "Somerset House East Wing, London",
			Postcode:     "WC2R 1LA",
			Longitude:    -0.1167651136,
			Latitude:     51.51151242,
			Condition:    "Make a purchase every 1.5 hours, some restrictions on laptop-use",
			ImgUrl:       "https://laptopfriendly.co/images/places/london/watchhouse-somerset-house/watchhouse-somerset-house%20other%20watchhouse-somerset-house-other%20.jpg",
			CreatedBy:    103,
		},
		{
			LocationName: "V&A Museum",
			Address:      "Cromwell Rd, London",
			Postcode:     "SW7 2RL",
			Longitude:    -0.172118583,
			Latitude:     51.49675506,
			Condition:    "No purchase needed",
			ImgUrl:       "https://hfemploymentandwellbeing.org.uk/wp-content/uploads/2022/04/v-a-museum.jpg",
			CreatedBy:    102,
		},
		{
			LocationName: "Senate House Library",
			Address:      "Senate House University of London, Malet St, London",
			Postcode:     "WC1E 7HU",
			Longitude:    -0.1292493519,
			Latitude:     51.52110574,
			Condition:    "Library membership required",
			ImgUrl:       "https://www.london.ac.uk/sites/default/files/styles/max_1300x1300/public/2017-09/_RMP2015_UOL_0216-13.jpg?itok=bybhltBf",
			CreatedBy:    102,
		},
		{
			LocationName: "Foyles",
			Address:      "107 Charing Cross Rd, London",
			Postcode:     "WC2H 0DT",
			Longitude:    -0.1297976136,
			Latitude:     51.51446172,
			Condition:    "Purchase required every 3 hours",
			ImgUrl:       "https://leafi.co.uk/sites/default/files/styles/slideshowbreakpoints_theme_fusion_starter_wide_1x/public/Foyles_picture_web.jpg?itok=r5ptyv7E",
			CreatedBy:    101,
		},
		{
			LocationName: "Lush Oxford Street",
			Address:      "175-179 Oxford St, London",
			Postcode:     "W1D 2JS",
			Longitude:    -0.1384737844,
			Latitude:     51.5156085,
			Condition:    "Make a purchase around every 2 hours",
			ImgUrl:       "https://www.oxfordstreet.co.uk/wp-content/uploads/fly-images/231997/lush_oxford_street_shop_2021_93_2-min-scaled-1300x923.jpg",
			CreatedBy:    105,
		},
		{
			LocationName: "Primark Cafe",
			Address:      "14-28 Oxford St, London",
			Postcode:     "W1D 1AU",
			Longitude:    -0.1311124001,
			Latitude:     51.51666801,
			Condition:    "Make a purchase every 3 hours",
			ImgUrl:       "https://www.videea.com/wp-content/uploads/2021/05/primark-london-10.jpg",
			CreatedBy:    104,
		},
		{
			LocationName: "Ozone Coffee",
			Address:      "11 Leonard St, London",
			Postcode:     "EC2A 4AQ",
			Longitude:    -0.08664137482,
			Latitude:     51.52523972,
			Condition:    "Make a purchase around every 2 hours",
			ImgUrl:       "https://4.bp.blogspot.com/-A3b0T1bQtOM/V089jtBmBgI/AAAAAAAAOqg/D6mejiCU_2kJrnfQ_I8dfMLH50f9Wd6TwCLcB/s1600/IMG_9505.JPG",
			CreatedBy:    103,
		},
		{
			LocationName: "District",
			Address:      "New Union Square Embassy Gardens, 7 Ponton Rd, Nine Elms, London",
			Postcode:     "SW11 7DN",
			Longitude:    -0.1337516939,
			Latitude:     51.48213786,
			Condition:    "Make a purchase every 2 hours or so",
			ImgUrl:       "https://images.squarespace-cdn.com/content/v1/5c406fa889c172970e5208e9/1570615000324-YAN8M180GLA2JREIHXYE/IMG_0018.jpg",
			CreatedBy:    102,
		},
		{
			LocationName: "The Common E2",
			Address:      "53 Old Bethnal Green Rd, Bethnal Green, London",
			Postcode:     "E2 6QA",
			Longitude:    -0.06065613452,
			Latitude:     51.53017239,
			Condition:    "Make a purchase every 2 hours",
			ImgUrl:       "https://media.timeout.com/images/103516880/image.jpg",
			CreatedBy:    101,
		},
	}

	LocationFeatures := []models.LocationFeature{
		{
			LocationId:           1,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       true,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           2,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                true,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          true,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                true,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         false,
		},
		{
			LocationId:           3,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 false,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          true,
			Vegetarian:           true,
			Vegan:                true,
			OutdoorSeating:       true,
			PhoneSignal:          false,
			NaturalLight:         false,
		},
		{
			LocationId:           4,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                true,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           4,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         false,
			Noisy:                true,
			Food:                 true,
			Beverages:            true,
			Toilets:              false,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           false,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         false,
		},
		{
			LocationId:           6,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           7,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                true,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           8,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                true,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          true,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         false,
		},
		{
			LocationId:           9,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         false,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           10,
			WiFi:                 false,
			WaterFountain:        false,
			PowerSockets:         false,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           11,
			WiFi:                 false,
			WaterFountain:        false,
			PowerSockets:         false,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           12,
			WiFi:                 true,
			WaterFountain:        true,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         false,
		},
		{
			LocationId:           13,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         false,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              false,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           14,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          true,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           15,
			WiFi:                 true,
			WaterFountain:        true,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          false,
			NaturalLight:         true,
		},
		{
			LocationId:           16,
			WiFi:                 true,
			WaterFountain:        true,
			PowerSockets:         true,
			Noisy:                true,
			Food:                 false,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           17,
			WiFi:                 true,
			WaterFountain:        true,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           18,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           19,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         false,
			Noisy:                true,
			Food:                 true,
			Beverages:            true,
			Toilets:              false,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           20,
			WiFi:                 true,
			WaterFountain:        false,
			PowerSockets:         true,
			Noisy:                true,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          false,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         false,
		},
		{
			LocationId:           21,
			WiFi:                 true,
			WaterFountain:        true,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          true,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           22,
			WiFi:                 true,
			WaterFountain:        true,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     true,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          true,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
		{
			LocationId:           23,
			WiFi:                 true,
			WaterFountain:        true,
			PowerSockets:         true,
			Noisy:                false,
			Food:                 true,
			Beverages:            true,
			Toilets:              true,
			WheelchairAccess:     false,
			DiasabledToilets:     false,
			GenderNeutralToilets: false,
			FreeParking:          false,
			PaidParking:          false,
			PetFriendly:          true,
			Vegetarian:           true,
			Vegan:                false,
			OutdoorSeating:       false,
			PhoneSignal:          true,
			NaturalLight:         true,
		},
	}

	for _, user := range Users {
		model.AddUser(&user)
	}

	for _, location := range Locations {
		model.AddLocation(&location)
	}

	for _, location := range LocationFeatures {
		model.AddLocationFeatures(&location)
	}
}
