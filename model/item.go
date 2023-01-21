package model

import (
	"time"

	"gorm.io/gorm"
)

// item 商品模型
type Item struct {
	gorm.Model
	Name       string
	Brief      string
	Picture    string
	Introduce  string
	Algorithm  string
	Code       string
	Time       int // 单位ms
	Memory     int // 单位b
	OutputImg  bool
	AuthorId   uint
	Status     int
	ReviewId   uint
	ReviewTime *time.Time
}

const (
	// 待审核
	Pending = 1
	// 通过
	Accept = 2
	// 拒绝
	Reject = 3

    TimeLayout = "2006-01-02 15:04:05"
)


// GetItem 用ID获取商品
func GetItem(ID interface{}) (Item, error) {
	var item Item
	result := DB.First(&item, ID)
	return item, result.Error
}
