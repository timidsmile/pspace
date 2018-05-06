package model

type BasicMdl struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreateTime uint64  `json:"createdTime"`
	UpdatedTime uint64  `json:"updatedTime"`
}