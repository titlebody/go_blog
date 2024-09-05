package model

import (
	"go_blog/global"
	"go_blog/model/c_type"
	"gorm.io/gorm"
	"os"
)

type BannerModel struct {
	MODEL
	Path string           `json:"path"`
	Hash string           `json:"hash"`
	Name string           `json:"name"`
	Type c_type.ImageType `gorm:"default:1" json:"image_type"` //图片类型，本地还是七牛
}

// 在同一个事务中更新数据
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.Type == c_type.Local {
		// 本地图片删除还要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error("删除本地图片失败", err)
			return err
		}
	}

	return nil
}
