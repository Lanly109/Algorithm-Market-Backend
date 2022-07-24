package model

import (
	"gorm.io/gorm"
)

// item 商品模型
type Item struct {
	gorm.Model
	Name      string
	Brief     string
	Picture   string
	Introduce string
	Algorithm string
	Code      string
	Time      int // 单位ms
	Memory    int // 单位kb
	OutputImg bool
}

// GetItem 用ID获取商品
func GetItem(ID interface{}) (Item, error) {
	var item Item
	result := DB.First(&item, ID)
	return item, result.Error
}
