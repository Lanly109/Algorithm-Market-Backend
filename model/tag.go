package model

import (
	"gorm.io/gorm"
)

// tag 标签模型
type Tag struct {
	gorm.Model
	TagName string
	ItemId  uint
}

// GetTags 用itemID获取标签
func GetTags(ID interface{}) ([]Tag, error) {
	var tags []Tag
	result := DB.Select("TagName").Where("item_id = ?", ID).Find(&tags)
	return tags, result.Error
}

// DeleteTags 用itemID删除标签
func DeleteTags(ID interface{}) error {
	err := DB.Where("item_id = ?", ID).Delete(&Tag{}).Error
	return err
}

// CreateTags 用itemID创建标签
func CreateTags(ID uint, tags []string) error {
	for _, tagName := range tags {
		tag := Tag{
			ItemId:  ID,
			TagName: tagName,
		}
		if err := DB.Create(&tag).Error; err != nil {
			return err
		}
	}
	return nil
}
