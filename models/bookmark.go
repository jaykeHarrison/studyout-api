package models

type Bookmark struct {
	UserRefer  Users     `gorm:"foreignKey:UserId"`
	UserId     uint    `gorm: not null; default:null`
	LocationId uint     `gorm: not null; default:null`
	Location   Location `gorm:"foreignKey:LocationId"`
}

