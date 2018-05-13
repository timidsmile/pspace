package model

type UserThirdBasic struct {
	Basic
	UserID         int64  `gorm:"column:user_id" json:"userID"`
	AuthType       uint8  `gorm:"column:user_name" json:"authType"`
	ThirdUserId    string `gorm:"column:third_user_id"`
	ThirdNickName  string `gorm:"column:third_nick_name"`
	ThirdAvatarUrl string `gorm:"column:third_avatar_url"`
	Status         uint8  `gorm:"default:1"` // 使用default关键字指定默认值
}

func (UserThirdBasic) TableName() string {
	return "user_third_basic"
}
