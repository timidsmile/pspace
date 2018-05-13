package model

type UserBasic struct {
	Basic
	UserID    int64 `gorm:"index;column:user_id;" json:"userID"`
	UserName  string `gorm:"index;size:32;column:user_name" json:"userName"`
	Mobile    string `gorm:"index;size:11" json:"mobile"`
	Email     string `gorm:"index;size:50" json:"email"`
	Passwd    string `gorm:"size:32" json:"passwd"`
	NickName  string `gorm:"size:50" json:"nickname"`
	AvatarUrl string `json:"avatarURL"`
	Status    int8   `gorm:"default:1" json:"status"` // 使用default关键字指定默认值

}

func (UserBasic) TableName() string {
	return "user_basic"
}
