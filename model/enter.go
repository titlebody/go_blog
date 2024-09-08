package model

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id" structs:"-"`
	CreatedAt time.Time `json:"created_at" structs:"-"`
	UpdatedAt time.Time `json:"updated_at" structs:"-"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
