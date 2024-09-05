package model

import "time"

type BannerModel struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	Path      string    `json:"path"`
	Hash      string    `json:"hash"`
	Name      string    `json:"name"`
}
