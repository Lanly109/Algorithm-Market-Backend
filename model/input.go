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
	ItemId uint
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

// DeleteInputs 用itemID删除输入样例
func DeleteInputs(ID interface{}) error {
	err := DB.Where("item_id = ?", ID).Delete(&Input{}).Error
	return err
}

// CreateInputs 用itemID创建输入样例
func CreateInputs(ID uint, inputs []string) error {
	for _, data := range inputs {
		input := Input{
			ItemId: ID,
			Input:  data,
		}
		if err := DB.Create(&input).Error; err != nil {
			return err
		}
	}
	return nil
}
