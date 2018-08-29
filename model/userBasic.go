package model

type User struct {
	Basic
	UserID    string `gorm:"size:32 column:user_id" json:"userID"`
	UserName  string `gorm:"size:32 column:user_name" json:"userName"`
	mobile    string `gorm:"size:32" json:"mobile"`
	email     string `gorm:"size:32" json:"email"`
	passwd    string `gorm:"size:32" json:"passwd"`
	Nickname  string `gorm:"size:32" json:"nickname"`
	AvatarURL string `gorm:"size:255" json:"avatarURL"`
	status    int8   `gorm:"size:255" json:"status"`
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
