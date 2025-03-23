package models

import "time"

type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

// PageInfo 表查询相关信息
type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

// RemoveRequest 删除列表
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
