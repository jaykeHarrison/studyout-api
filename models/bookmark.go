package models

type Bookmark struct {

	UserRefer  uint
	UserId  Users `gorm:"foreignKey:UserRefer"`

	LocationId    uint
	Location      Location `gorm:"foreignKey:LocationId"`
	
}