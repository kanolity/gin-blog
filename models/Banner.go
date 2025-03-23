package models

import (
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models/ctype"
	"gorm.io/gorm"
	"os"
)

type Banner struct {
	Model
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片存储位置
}

func (b *Banner) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//本地图片删除要删除本地的存储
		err := os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
