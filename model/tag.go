package model

import (
	"gorm.io/gorm"
)

// tag 标签模型
type Tag struct {
	gorm.Model
	TagName        string
	ItemId         int
}

// GetTags 用itemID获取标签
func GetTags(ID interface{}) ([]Tag, error) {
	var tags []Tag
	result := DB.Select("TagName").Where("item_id = ?", ID).Find(&tags)
	return tags, result.Error
}
