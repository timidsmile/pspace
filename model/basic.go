package model

type Basic struct {
	// 列名为字段名的蛇形小写
	ID         uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id"` // 列名为 `id`, 也可以用 column显示指定列名; 字段`ID`为默认主键,也可使用primary_key指定主键
	CreateTime uint64 `json:"createTime"`                           // 列名为 `create_time`
	UpdateTime uint64 `json:"updateTime"`
}
