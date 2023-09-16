package models

// 点赞信息
type Favorite struct {
	ID      uint  `gorm:"primarykey; comment:点赞id"`
	UserId  uint  `gorm:"not null; comment:用户ID;  type:INT"`
	VideoId uint  `gorm:"not null; comment:视频ID;  type:INT"`
	User    User  `gorm:"foreignKey:UserId;  references:ID; comment:点赞用户的信息"`
	Video   Video `gorm:"foreignKey:VideoId; references:ID; comment:点赞视频的信息"`
}
