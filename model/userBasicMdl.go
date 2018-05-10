package model

type UserBasic struct {
	BasicMdl
	UserID              string `gorm:"size:32 column:user_id" json:"userID"`
	UserName              string `gorm:"size:32 column:user_name" json:"userName"`
	Mobile              string `gorm:"size:32" json:"mobile"`
	Email              string `gorm:"size:32" json:"email"`
	Passwd              string `gorm:"size:32" json:"passwd"`
	NickName          string `gorm:"size:32" json:"nickname"`
	AvataUrl         string `gorm:"size:255" json:"avatarURL"` // TODO: AvatarUrl
	Status         int8 `gorm:"size:255" json:"status"`

}

/*
| id          | bigint(64) unsigned | NO   | PRI | NULL    | auto_increment |
| user_id     | bigint(20)          | NO   | MUL | NULL    |                |
| user_name   | varchar(32)         | NO   | MUL |         |                |
| mobile      | varchar(11)         | NO   |     |         |                |
| email       | varchar(50)         | NO   | MUL |         |                |
| passwd      | varchar(32)         | NO   |     |         |                |
| nick_name   | varchar(50)         | NO   |     |         |                |
| avata_url   | text                | YES  |     | NULL    |                |
| create_time | int(10) unsigned    | NO   |     | NULL    |                |
| update_time | int(10) unsigned    | NO   |     | NULL    |                |
| status      | s
 */


func (UserBasic) TableName() string {
	return "user_basic"
}