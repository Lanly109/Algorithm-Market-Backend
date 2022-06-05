package model

import (
	"gorm.io/gorm"
)

// Output 输出模型
type Output struct {
	Type string
	Data string
}

// Input 输入模型
type Input struct {
	gorm.Model
	Input  string
	Output Output `gorm:"embedded;embeddedPrefix:Output_"`
	ItemId int
}

// GetInputs 用itemID获取输入
func GetInputs(ID interface{}) ([]Input, error) {
	var inputs []Input
	result := DB.Where("item_id = ?", ID).Find(&inputs)
	return inputs, result.Error
}

// GetInput 用ID获取输入
func GetInput(ID interface{}) (Input, error) {
	var input Input
	result := DB.Find(&input, ID)
	return input, result.Error
}
