package user

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(255)"`
	Password string `gorm:"column:password;type:varchar(255);"json:"password"`
	// 关注数与粉丝数
	FollowCount    int    `gorm:"column:follow_count;type:bigint(20)"`
	FollowerCount  int    `gorm:"column:follower_count;type:bigint(20)"`
	TotalFavorited int    `gorm:"column:total_favorited;type:bigint(20)"`
	FavoriteCount  int    `gorm:"column:favorite_count;type:bigint(20)"`
	Avatar         string `gorm:"column:avatar;type:varchar(255)"`
	BgImage        string `gorm:"column:bg_image;type:varchar(255)"`
	Signature      string `gorm:"column:signature;type:varchar(255)"`
	// 已关注用户Id
	//FollowerID string
	// 权限，默认为1
	//Authority int `gorm:"default:1"`
}

func (u Users) TableName() string {
	return "users"
}
