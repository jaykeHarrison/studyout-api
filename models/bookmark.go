package models

type Bookmark struct {
	UserRefer  uint     `gorm: not null; default:null`
	UserId     Users    `gorm:"foreignKey:UserRefer"`
	LocationId uint     `gorm: not null; default:null`
	Location   Location `gorm:"foreignKey:LocationId"`
}
