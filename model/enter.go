package model

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
